package sync

import (
	"sync"
	"fmt"
	"testing"
)

/**
 *我们可以利用sync.WaitGroup来满足这样的情况：
  ▲某个地方需要创建多个goroutine，并且一定要等它们都执行完毕后再继续执行接下来的操作。
  是的，WaitGroup最大的优点就是.Wait()可以阻塞到队列中的任务都完毕后才解除阻塞。
 */
var waitGroup sync.WaitGroup

func Call(shutDown int){
	fmt.Println(shutDown)
	waitGroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}



func Test_Call(t *testing.T) {
	for i := 0; i<10; i++{
		waitGroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Call(i)
	}
	waitGroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}