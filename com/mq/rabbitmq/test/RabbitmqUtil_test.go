package rabbitmq

import (
	"testing"
	"golang-test/com/mq/rabbitmq/service"
	"fmt"
	"time"
	"runtime"
	"golang-test/com/mq/rabbitmq/constants"
)

func Test_receive(t *testing.T) {

	go func() {
		for {
			go service.Push(constants.TL)
			time.Sleep(1 * time.Second)
		}
	}()
	service.Receive()
	fmt.Println("end")
	service.Close()
}

func Test_Call(t *testing.T) {

	go func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
		for {
			var limit = constants.THOUSAND
			channel := make(chan int, limit)
			for i := 0; i < limit; i++ {
				go func() {
					channel <- service.Push(constants.TL)
				}()
				fmt.Println(<-channel)
			}

			close(channel)
		}

		//for {
		//	go service.Push("TL")
		//	time.Sleep(1 * time.Second)
		//}
	}()
	service.Receive()
	fmt.Println("end")
	service.Close()
}

func Test_Receive(t *testing.T) {

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