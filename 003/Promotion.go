package main

import (
	"fmt"
)

//Person is Exported type
type Person struct {
	FirstName string
	LastName  string
	age       int //age isn't accessible directly.
}

//Customer is exported type
type Customer struct {
	Person    //embedded type, Customer tyope can call Person fields directly
	ContactNo int
}

func (per Person) showAge() int {
	return per.age
}

func (per *Person) show() {
	fmt.Println(per.FirstName, "", per.LastName, "", per.showAge())
}
func (cust *Customer) show() { //method overriding based on receiver
	cust.Person.show() // option1: Calling Fields from Person struct and then the ContactNo from Customer
	fmt.Println(cust.ContactNo)
	fmt.Println(cust.FirstName, "", cust.LastName, "", cust.showAge(), cust.ContactNo) //option2: accessing embeded type fields directly
}

func main() {
	cust := Customer{
		Person: Person{
			"John",
			"Matthew",
			25,
		},
		ContactNo: 9965,
	}

	individual1 := Person{
		"Mike",
		"Johnsson",
		37,
	}

	customer1 := Customer{
		individual1,
		1234,
	}
	cust.show()
	customer1.show()
}
