package main

import (
	"fmt"
	"time"
)

func main() {
	chanmsg := make(chan int)
	done := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			chanmsg <- i
		}
		close(chanmsg)
		close(done)
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
				case <-done:
					close(next)
					return
				}
			}
		}()
		return next
	}

	a := recv(recv(chanmsg, 1), 2)
	for msg := range a {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("3 got message", msg)
	}

}
