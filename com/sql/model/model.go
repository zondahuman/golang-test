package model

import "time"

type OrderInfo struct {
	Id int
	Name string
	Age  int
	CreateTime time.Time
	UpdateTime time.Time
	Version int
}

type OrderInfoResponse struct {
	Id int64
	Affect  int64

}
