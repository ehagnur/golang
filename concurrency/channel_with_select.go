//package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every second"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "Every five seconds"
			time.Sleep(time.Second * 5)
		}
	}()

  //The 5 second interval blocks the 1 second
  //for {
  //  fmt.Println(<-c1)
  //  fmt.Println(<-c2)
  //}

  //select will print which ever is ready. with out select the two second will block the 500ms call
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
