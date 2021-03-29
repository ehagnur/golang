package main

import "fmt"

// Employee struct type
// There is no concept like public, private, protected for the visibility purpose in Golang.
// As Golang is simple, if the first letter of the struct name or variable name is capital
// then it is visible outside that package. check 003, Promotion.go
type Employee struct {
	employeeid int
	firstname  string
	lastname   string
	deptid     int
}

func (emp Employee) getID() int {
	return emp.employeeid
}

func (emp *Employee) setDeptID(ID int) {
	emp.deptid = ID
}

func main() {
	e1 := Employee{ //instanciates a new Employee
		employeeid: 100,
		firstname:  "Hagos",
		lastname:   "Salih",
		deptid:     1,
	}

	e2 := new(Employee) //allocate memory to all fields by setting default values and returning pointer to it (*Employee)
	e2.setDeptID(99)
	fmt.Printf("e2 default employee id is: %d and dept id: %d \n",
		e2.getID(), e2.deptid)
	e2.employeeid = 101
	fmt.Printf("e2 employee id is set to: %d\n", e2.getID())

	//fmt.Println(employee.getID())
	fmt.Println(e1)
	e1.setDeptID(10)
	fmt.Println(e1)

	if e1.firstname == "Hagos" {
		fmt.Printf("Employee name is %s\n", e1.firstname)
	} else {
		fmt.Printf("Nope, this is different %s\n", e1.firstname)
	}

}
