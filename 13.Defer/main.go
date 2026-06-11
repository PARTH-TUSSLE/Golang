package main

import (
	"errors"
	"fmt"
)

func main() {
	// makes sure cleanup happens everytime
	fmt.Println("Case 1:- Success")
	if err := doWork(true); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Case 2:- Failure")
	if err := doWork(false); err != nil {
		fmt.Println("Error", err)
	}
}

func doWork(success bool) error {
	// some resource realted work
	// start message -> resource acuired
	// cleanup message -> resource released

	fmt.Println("Start: resource acuired !",)

	// makes sure this will run at the end of the function , gonna run in both the cases
	defer fmt.Println("Cleanup: resource released !")
	
	if !success {
		return errors.New("Something went wrong, returning early")
	} 

	fmt.Println("Doing the work")
	fmt.Println("Work is done")

	return nil

}