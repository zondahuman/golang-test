package rabbitmq

import (
	"testing"
	"encoding/json"
	"fmt"
	"golang-test/com/basic/util/uuid"
	"golang-test/com/basic/common"
	"strconv"
	"golang-test/com/basic/util/random"
	"golang-test/com/mq/rabbitmq/service"
)

func Test_sendRabbitmq(t *testing.T) {
	source := "YM"
	service.SendRabbitmq(source)
}

func Test_sendApplication(t *testing.T) {
	source := "RRD"
	service.SendRabbitmq(source)
}

func Test_sendApply(t *testing.T) {
	source := "TL"
	service.SendRabbitmq(source)
}

func Test_sendApply1(t *testing.T) {
	channel := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go func() {
			channel <- service.SendLend("YM")
		}()
		fmt.Println(<-channel)
	}

	close(channel)
}

func Test_sendApply2(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " --start")
		go service.SendLend("YM")
		fmt.Println(i, " --end")

	}
}


func Test_send(t *testing.T) {
	str := "{ \"id\": \"0001\", \"type\": \"donut\", \"name\": \"Cake\", \"ppu\": 0.55, \"batters\": { " +
		"\"batter\": [{ \"id\": \"1001\", \"type\": \"Regular\" }, { \"id\": \"1002\", \"type\": " +
		"\"Chocolate\" }] }, \"topping\": [{ \"id\": \"5001\", \"type\": \"None\" }, { \"id\": \"5002\", " +
		"\"type\": \"Glazed\" }, { \"id\": \"5005\", \"type\": \"Sugar\" }, { \"id\": \"5007\", \"type\": " +
		"\"Powdered Sugar\" }] }"
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		panic(err)
	}
	batters := jsonMap["batters"].(map[string]interface{})["batter"]
	for _, b := range (batters.([]interface{})) {
		b.(map[string]interface{})["id"] = "other id"
	}
	str2, err := json.Marshal(jsonMap)
	fmt.Println(string(str2))
}

func Test_send1(t *testing.T) {

	fmt.Println(string(uuid.GetGuid()))
	fmt.Println(util.Number())
	userKey := uuid.GetGuid()
	fmt.Println("userKey=", userKey)
	result := common.Substr(userKey, 0, 15)
	fmt.Println("result=", result)

}

func Test_switchCase(t *testing.T) {
	var userKey string
	source := "YM"
	switch source {
	case "RRD":
		userKey = uuid.GetGuid()
	case "YM":
		userKey = strconv.FormatInt(util.TimestampNanos(), 10)
	default : /* Optional */
		userKey = ""
	}
	fmt.Println("userKey=", userKey)

}