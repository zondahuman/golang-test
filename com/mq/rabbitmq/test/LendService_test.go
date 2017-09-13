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
	"golang-test/com/mq/rabbitmq/constants"
	"time"
	"golang-test/com/basic/util/date"
)


func Test_lendLoanConfirmed(t *testing.T) {
	source := constants.CONFIRMED
	applicationId := "RRD_1969138wCc0yxoT3_1"
	service.LendMessageConfirm(source, applicationId)
}

func Test_lendLoanExpired(t *testing.T) {
	source := constants.EXPIRED
	applicationId := "RRD_b56b33637b32198_9"
	service.LendMessageConfirm(source, applicationId)
}

func Test_lendWebapp(t *testing.T) {
	source := constants.WEBAPP
	service.LendMessage(source)
}

func Test_lendBeauty(t *testing.T) {
	source := constants.YM
	service.LendMessage(source)
}

func Test_lendLoan(t *testing.T) {
	source := constants.RRD
	service.LendMessage(source)
}

func Test_lendLoanCredit(t *testing.T) {
	source := constants.RRDCREDIT
	service.LendMessage(source)
}

func Test_lendLoanWebapp(t *testing.T) {
	source := constants.RRDWEBAPP
	service.LendMessage(source)
}

func Test_lendLoanNullConfirm(t *testing.T) {
	source := constants.RRDN
	service.LendMessage(source)
}


func Test_lendTraining(t *testing.T) {
	source := constants.TL
	service.LendMessage(source)
}

func Test_lendFuse(t *testing.T) {
	source := constants.FU
	service.LendMessage(source)
}

func Test_lendLoanLoanTl1(t *testing.T) {
	source := constants.LOANTL
	service.LendMessage(source)
}

func Test_lendBatchFuse1(t *testing.T) {
	var limit = constants.FIVE
	channel := make(chan int, limit)
	for i := 0; i < limit; i++ {
		go func() {
			channel <- service.LendMessage(constants.FU)
		}()
		fmt.Println(<-channel)
	}

	close(channel)
}

func Test_lendBatchTl1(t *testing.T) {
	var limit = constants.FIVE
	channel := make(chan int, limit)
	for i := 0; i < limit; i++ {
		go func() {
			channel <- service.LendMessage(constants.TL)
		}()
		fmt.Println(<-channel)
	}

	close(channel)
}

func Test_lendBatchRrd(t *testing.T) {
	var limit = constants.FIVE
	channel := make(chan int, limit)
	for i := 0; i < limit; i++ {
		go func() {
			channel <- service.LendMessage(constants.RRD)
		}()
		fmt.Println(<-channel)
	}

	close(channel)
}

func Test_lendBatchLoanTL(t *testing.T) {
	var limit = constants.FIVE
	channel := make(chan int, limit)
	for i := 0; i < limit; i++ {
		go func() {
			channel <- service.LendMessage(constants.LOANTL)
		}()
		fmt.Println(<-channel)
	}

	close(channel)
}



func Test_sendLendApply(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " --start")
		go service.LendMessage(constants.YM)
		fmt.Println(i, " --end")

	}
}


func Test_sendLend1(t *testing.T) {
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

func Test_sendLend(t *testing.T) {

	fmt.Println(string(uuid.GetGuid()))
	fmt.Println(util.Number())
	userKey := uuid.GetGuid()
	fmt.Println("userKey=", userKey)
	result := common.Substr(userKey, 0, 15)
	fmt.Println("result=", result)

}

func Test_switchLendCase(t *testing.T) {
	var userKey string
	source := constants.YM
	switch source {
	case constants.RRD:
		userKey = uuid.GetGuid()
	case constants.YM:
		userKey = strconv.FormatInt(util.TimestampNanos(), 10)
	case constants.TL:
		userKey = uuid.GetGuid()
	default : /* Optional */
		userKey = ""
	}
	fmt.Println("userKey=", userKey)

}


func Test_sendLendId(t *testing.T) {
	result := service.SpliceLendId("RRD_4efe866496ce919_9")
	fmt.Println("result=", result)
}


func Test_sendTimestamp(t *testing.T) {
	result := date.CurrentTimestamp()
	fmt.Println("result=", result)
	cur := time.Now()
	timestamp := cur.UnixNano() / 1000000
	fmt.Println("timestamp=", timestamp)
}


