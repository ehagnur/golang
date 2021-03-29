package main

import "fmt"

func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func main() {
	//Varidic fucntion take arbitrary number of arduments.
	//println function is an example of variadic function
	fmt.Println("This", "is", "an", "example")

	//call mult function
	fmt.Println(mult(1, 2, 3, 4))

	//variadic function can be applied to slices
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(mult(nums...))
}
