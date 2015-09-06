// goroutine.go
package main

import (
	"fmt"
	"runtime"
)

func say(s string){
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func goSay() {
	go say("world") //start new goruntine
	say("hello") //current goruntine
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum //send sum to c
}

func fibonacci(n int, c chan int){
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
		//fmt.Println(x, y)
	}
	close(c)
}

func invokeChan() {
	a := []int{7,8,9,6,5,2}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	
	fmt.Println(x, y, x + y)
	
	c2 := make(chan int, 3)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2,<-c2)
	
	
}





func main() {
	//invokeChan()
	c3 := make(chan int, 10)
	fmt.Println(cap(c3))
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}
}



