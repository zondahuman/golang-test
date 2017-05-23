package redis

import (
	"testing"
)

func Test_operateSetnx(t *testing.T) {
	//operateSetnx("abin", "lee")
	num := operateSetnx("abin1", "lee")
	t.Log("num=", num)
}

func Test_operateSet(t *testing.T) {
	operateSet("abin", "lee")
	//num := operateSet("abin1", "lee")
	//t.Log("num=", num)
}


func Test_operateGet(t *testing.T) {
	operateGet("abin")
	//num := operateGet("abin1", "lee")
	//t.Log("num=", num)
}