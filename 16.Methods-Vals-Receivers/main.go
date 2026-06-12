package main

import(
	"fmt"
)

type User struct{
	name string
	age int
}

func main() {
	u := User{name: "Parth", age: 21}
	fmt.Println("Before Intro", u.age)
	u.Intro()
	fmt.Println("After Intro", u.age)
	fmt.Println("Before", u.age)
	u.Birthday()
	fmt.Println("After", u.age)
}

// this method will receive a copy of the user 
func (u User) Intro() {  // function attached to a type
	fmt.Println("Hello my name is", u.name)

	u.age++

}


// Methods-pointer-receiver



type User2 struct{
	name string
	age int
}

// this method will receive a copy of the user 
func (u *User) Birthday() {  // function attached to a type
	u.age++
}