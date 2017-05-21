package model

import "time"

type OrderInfo struct {
	id int
	name string
	age  int
	createTime time.Time
	updateTime time.Time
	version int
}

type OrderInfoResponse struct {
	Id int64
	Affect  int64

}
