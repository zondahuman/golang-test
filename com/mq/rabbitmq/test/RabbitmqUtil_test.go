package rabbitmq

import (
	"testing"
	"golang-test/com/mq/rabbitmq/service"
	"fmt"
	"time"
)

func Test_receive(t *testing.T) {

	go func() {
		for {
			service.Push("TL")
			time.Sleep(1 * time.Second)
		}
	}()
	service.Receive()
	fmt.Println("end")
	service.Close()
}

