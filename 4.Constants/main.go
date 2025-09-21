package main

import "fmt"

func main() {
	const name = "Golang"

	// name = "JavaScript"

	fmt.Println(name)

	const  (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}