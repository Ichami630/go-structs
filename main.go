package main

import "fmt"

/*
when declaring a struct, the name of the struct has a lot to tell about the scope the struct
lowercase means the struct has a private scope
uppercase means the struct has a public scope
same applies to the variables or fields in the struct

we declare a struct using the type keyword followed by the name of the struct and the struct keyword
*/
type Employee struct {
	name     string
	age      int
	isRemote bool
	Address  //embedded struct
}

/*
Embedded structs (Inheritance in Go)
Go doesn't have class inheritance, but you can embed one struct inside another
*/
type Address struct {
	Street string
	City   string
}

// func to update the name in the struct
func (e *Employee) updateName(newName string) {
	e.name = newName
}

func main() {
	address := Address{
		Street: "123 mayor street",
		City:   "Buea",
	}
	employee1 := Employee{
		name:     "ichami",
		age:      10,
		isRemote: false,
		Address:  address,
	}
	employee1.updateName("brandon")
	fmt.Println("Employee name ", employee1.name, " is ", employee1.age, " and remote working is ", employee1.isRemote)
	fmt.Println("an embedded struct fields:", employee1.Street) //accessing embedded fields

	//declaring an anonymous struct
	job := struct {
		title  string
		salary int
	}{
		title:  "fullstack",
		salary: 900,
	}

	fmt.Println("job:", job.title, "salary:", job.salary)
}
