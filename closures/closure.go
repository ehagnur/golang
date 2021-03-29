package main

import "fmt"

//returnAnonFunc returns annonymous function
func returnAnonFunc() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

//a closure function value that references local variables in the scope it is defined
//intSeq returns a function that returns int
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func appendTag() func(string) string {
	tags := ""
	return func(tag string) string {
		tags += tag + " "
		return tags
	}

}

func main() {

	tagAppender := appendTag()
	fmt.Println(tagAppender("SFROLE=XP"))
	tagAppender("SFDEPLOYMENTGROUP=prod")
	fmt.Println(tagAppender("SFCLOUD=aws"))

	/*
		  //annonymous function
			func(msg string) {
				fmt.Println(msg)
			}("Hello Annonymous")

			//returned anonnymous function is assigned to a variable
			printFunc := returnAnonFunc()
			printFunc("Hello return from anonnymous") //running the annonymous function

			nextInt := intSeq()
			fmt.Println(nextInt()) //prints 1
			fmt.Println(nextInt()) //prints 2: nextInt keeps track of the variable value i

			anotherInt := intSeq()    //A separate instances of the function
			fmt.Println(anotherInt()) //Prints 1
	*/
}
