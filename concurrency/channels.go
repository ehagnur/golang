//package main

import (
	"fmt"
	"time"
)


//Channels are used to send messages between routines
//Channels have type
//sending and receving messages is a blocking call

func main() {

		//Simple message passing
		c := make(chan string)
		go countWithChannel("sheep", c)

		for msg := range c {
			fmt.Println(msg)
		}
}


func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing //sending thing to Channel c
		time.Sleep(time.Microsecond * 1000000)
	}
	close(c)
}
