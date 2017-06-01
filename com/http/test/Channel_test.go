package http

import (
	"testing"
	"fmt"
)

func Test_Chan(t *testing.T) {
	fmt.Println("aa")
	chanel := make(chan int, 2)
	chanel <- 1
	chanel <- 2
	chanel <- 3
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)

}
func Test_Chan1(t *testing.T) {
	fmt.Println("Begin doing something!")
	channel := make(chan bool)
	go func(){
		fmt.Println("Doing Something..")
		close(channel)
	}()
	<-channel
	fmt.Println("Done..")

}
func worker(start chan bool, index int) {
        <-start
        fmt.Println("This is Worker:", index)
}
func Test_Chan2(t *testing.T) {
	 start := make(chan bool)
        for i := 1; i <= 100; i++ {
                go worker(start, i)
        }
        close(start)
        select {} //deadlock we expected

}