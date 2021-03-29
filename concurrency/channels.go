/*package main

import (
	"fmt"
	"time"
)


Channels are used to send messages between routines
Channels have type
sending and receving messages is a blocking call

func main() {
	/*
		//Simple message passing
		c := make(chan string)
		go countWithChannel("sheep", c)

		for msg := range c {
			fmt.Println(msg)
		}

	//select example
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

	for {
		select { //select will print which ever is ready. with out select the two second will block the 500ms call
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

/*
func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing //sending thing to Channel c
		time.Sleep(time.Microsecond * 500)
	}
	close(c)
} */
