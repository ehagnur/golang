package main

import "fmt"

func main() {
	// Array declation and creattion
	// var var_name [size]var_type
	var x [5]int //declation
	x[3] = 1000
	fmt.Println(x)
	b := [5]int{54, 76, 34, 2, 3} //declare and intialize

	fmt.Println(b) //print the entire Array

	for i, n := range b { //access index and value
		fmt.Printf("Index: %d => Value: %d\n", i, n)
	}

	// slice declation and creattion
	// var var_name []var_type
	var nums []int //no need to specify length
	fmt.Println(nums)
	nums = append(nums, 2, 7, 9)
	fmt.Println(nums)

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)

	// map declation and creattion
	//map [keyType] valueType
	var dictionary = make(map[string]int)
	dictionary["Zero"] = 999
	fmt.Println(dictionary["Zero"]) //Accessing value using key

	states := make(map[string]string)
	states["California"] = "CA"
	states["Texas"] = "TX"
	states["Florida"] = "FL"

	for k, v := range states {
		fmt.Printf("%s is the short name for %s\n", v, k)
	}

	for _, v := range states { //if the key or value is not needed use _ instead
		fmt.Println(v)
	}

}
