package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 40)
	results := make(chan int, 40)

	//go worker(jobs, results)
	//go worker(jobs, results)
	//go worker(jobs, results)
	//go worker(jobs, results)

	for i := 0; i < 40; i++ {
		jobs <- i
	}
	close(jobs)

	for k := 0; k < 4; k++ {
		go worker(k, jobs, results)
	}

	for i := 0; i < 40; i++ {
		fmt.Println(<-results)
	}

}

func worker(k int, jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Printf("Routine %d picking %d from jobs queue\n", k, n)
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
