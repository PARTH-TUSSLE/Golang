package main

import "fmt"

//maps -> hash tables, objects, dicts

func main() {
	//creation
	m := make(map[string]string)

	m["name"] = "Parth"
	fmt.Println(m["name"], len(m)) // if key DNE -> returns zeroes value of that data type(of the value)
	
	delete(m, "name");
	fmt.Println(m["name"], len(m))

	clear(m) // -> deletes all the elements in the map

	m2 := map[string]int { "price": 200, "age": 12 }

	fmt.Println(m2)

	// element exists or not

	v, ok := m2["price"]
	if ok {
		fmt.Println("All ok!", v)
	} else {
		fmt.Println("not ok!", v)
	}

}