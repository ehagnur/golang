package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 40)
	results := make(chan int, 40)

	for i := 0; i < 40; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 4; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 40; i++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

/*var wg sync.WaitGroup
		wg.Add(2) //telling the wait group there are two routine to wait for

		go func() { //precceding a function call with go makes it a routine
			count("one", 5)
			wg.Done()
		}()

		go func() {
			count("two", 40)
			wg.Done()
		}()

		for i := 5; i < 7; i++ {
			wg.Add(1)
			go func() {
				count(strconv.Itoa(i), i)
				wg.Done()
			}()
		}

		wg.Wait() //main routine should wait for the above two routines to finish
	}

	func count(thing string, freq int) {
		for i := 1; i <= freq; i++ {
			fmt.Println(i, thing)
			time.Sleep(time.Millisecond * 400)
		}
}
*/
