package main

import (
	"examples/imports/controllers"
	"fmt"
	"time"
)

func main() {
	fmt.Println(controllers.CheckLogin())
	fmt.Println("Checking page ...")
	time.Sleep(500)
	fmt.Println(controllers.CheckPageLoading())
}
