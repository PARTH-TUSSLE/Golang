package main
import (
	"fmt"
)

type User struct {
	id int
	name string
	email string
	age int
}

func main() {
	u1 := User{
		id: 124124,
		name: "Parth",
		email: "parth@gmailcom",
		age: 21,
	}

	u2 := User{
		name: "Parth",
		email: "abc@gmail.com",
	}

	fmt.Println(u1, u1.name)
	fmt.Println("Partial User ", u2)
}