package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go doesn't use exceptions for normal failures
	// functions return errors as normal return values

	// value, err := something()
	// if err != nil { handle the errro }

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	input := "30"

	level, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("Selected level:-", level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value, err) -> return pattern
	// nil error -> success
	// !nil -> error/failure
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Must be a number")
	}
	if n < 1 || n > 5 {
		return 0, fmt.Errorf("Must be in between 1 and 5")
	}
	return n, nil
}