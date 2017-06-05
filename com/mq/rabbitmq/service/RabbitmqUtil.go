package service

import (
	"fmt"
	"log"
	"bytes"
	"time"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel
var count = 0

const (
	queueName = "audit.request.queue"
	routingKey = "auditRequestKey"
	queueResponseName = "audit.response.tl.queue"
	exchange = "audit.exchange"
	mqurl = "amqp://guest:guest@172.16.2.145:15671/"
	//mqurl ="amqp://guest:guest@172.16.2.145:15671"
)

func main() {
	go func() {
		for {
			Push("TL")
			time.Sleep(1 * time.Second)
		}
	}()
	Receive()
	fmt.Println("end")
	Close()
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func Close() {
	channel.Close()
	conn.Close()
}

//连接rabbitmq server
func Push(source string) int {

	if channel == nil {
		mqConnect()
	}
	msgContent := splice(source)

	channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		MessageId:   string(1),
		Type:        "",
		Body:        []byte(msgContent),
	})
	return 0
}

func Receive() {
	if channel == nil {
		mqConnect()
	}

	msgs, err := channel.Consume(queueResponseName, "", true, false, false, false, nil)
	failOnErr(err, "")

	forever := make(chan bool)

	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			s := BytesToString(&(d.Body))
			count++
			fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}