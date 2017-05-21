package mysql

import "testing"
import "golang-test/com/sql/model"





func Test_findByPk(t *testing.T) {
	num := findByPk(11)
	t.Log("num=", num)
}

func Test_insert(t *testing.T) {
	//var orderInfo *model.OrderInfoResponse = new(model.OrderInfoResponse)
	orderInfo := insert("abin1", 56).(*model.OrderInfoResponse)
	t.Log("id=", orderInfo.Id, ", affect=", orderInfo.Affect)
}

func Test_findParam(t *testing.T) {
	num := findParam(11)
	t.Log("num=", num)
}