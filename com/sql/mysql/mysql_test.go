package mysql

import "testing"
import (
	"golang-test/com/sql/model"
	"encoding/json"
)





func Test_findByPk(t *testing.T) {
	num := findByPk(11)
	t.Log("num=", num)
}

func Test_insert(t *testing.T) {
	//var orderInfo *model.OrderInfoResponse = new(model.OrderInfoResponse)
	//orderInfoResponse := insert("abin1", 56)
	orderInfoResponse := insert("abin1", 56).(*model.OrderInfoResponse)
	t.Log("id=", orderInfoResponse.Id, ", affect=", orderInfoResponse.Affect)
	json.Marshal(orderInfoResponse)
}

func Test_findParam(t *testing.T) {
	orderInfo := findParam(11)
	//t.Log("orderInfo=", orderInfo)
	json.Marshal(orderInfo)
}