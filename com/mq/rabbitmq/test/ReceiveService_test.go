package rabbitmq

import (
	"fmt"
	"runtime"
	"time"
	"golang-test/com/mq/rabbitmq/service"
	"testing"
)

func init() {
	fmt.Println("init------start")
	go func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
		for {
			service.Receive()
			time.Sleep(1 * time.Second)
		}

		//for {
		//	go service.Push("TL")
		//	time.Sleep(1 * time.Second)
		//}
	}()
	fmt.Println("init------end")
}


func Test_Consumer(t *testing.T) {

	go func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
		for {
			service.Receive()
			time.Sleep(1 * time.Second)
		}

		//for {
		//	go service.Push("TL")
		//	time.Sleep(1 * time.Second)
		//}
	}()
	fmt.Println("end")
}