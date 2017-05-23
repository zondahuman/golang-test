package util

import (
	"time"
	"math/rand"
)

func Number100() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(100)
}

func Number10() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(10)
}

func Number() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Int()
}

func Number64() int64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Int63()
}

func TimestampSec() int64 {
	t := time.Now() //获取当前时间的结构体
	secs := t.Unix()          // 秒
	return secs
}

func TimestampNanos() int64 {
	t := time.Now() //获取当前时间的结构体
	nanos := t.UnixNano()     // 纳秒
	return nanos
}

