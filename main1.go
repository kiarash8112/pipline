package main

import (
	"fmt"
	"time"
)

func main() {
	chanmsg := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			chanmsg <- i
		}
	}()

	recv := func(chanmsg chan int, id int) chan int {
		next := make(chan int)
		go func() {
			for {
				select {
				case b := <-chanmsg:
					time.Sleep(time.Millisecond * 200)
					fmt.Println(id, "got message", b)
					next <- b
				}
			}
		}()
		return next
	}

	lastReciver := func(chanmsg chan int, id int) {
		go func() {
			for {
				select {
				case b := <-chanmsg:
					time.Sleep(time.Millisecond * 200)
					fmt.Println(id, "got message", b)
				}
			}
		}()
	}

	lastReciver(recv(recv(chanmsg, 1), 2), 3)

	time.Sleep(time.Second * 5)
}
