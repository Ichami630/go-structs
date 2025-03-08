package main

import (
	"encoding/json"
	"fmt"
)

/*
when declaring a struct, the name of the struct has a lot to tell about the scope the struct
lowercase means the struct has a private scope
uppercase means the struct has a public scope
same applies to the variables or fields in the struct

we declare a struct using the type keyword followed by the name of the struct and the struct keyword
*/

// structs with JSON tags
type Employee struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"isRemote"`
	Address         //embedded struct
}

/*
Embedded structs (Inheritance in Go)
Go doesn't have class inheritance, but you can embed one struct inside another
*/
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

// func to update the name in the struct
func (e *Employee) updateName(newName string) {
	e.Name = newName
}

// struct with JSON tags
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	address := Address{
		Street: "123 mayor street",
		City:   "Buea",
	}
	employee1 := Employee{
		Name:     "ichami",
		Age:      10,
		IsRemote: false,
		Address:  address,
	}
	employee1.updateName("brandon")
	fmt.Println("Employee name ", employee1.Name, " is ", employee1.Age, " and remote working is ", employee1.IsRemote)
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

	//tags
	jsonData, _ := json.Marshal(employee1)
	// fmt.Println(jsonData) //returns the encoded data as bytes
	fmt.Println(string(jsonData)) //returns the encoded data as json {"name":"brandon","age":10,"isRemote":false,"street":"123 mayor street","city":"Buea"}

	//json deserialisation
	//note previousely, we had a struct where we serialised the struct data and printed the json object
	//now we will do the reverse

	jsonString := `{"username":"ichami","email":"brandonichami@gmail.com","age":10}`

	var user User
	json.Unmarshal([]byte(jsonString), &user)
	fmt.Println(user.Username, user.Email, user.Age)

}
