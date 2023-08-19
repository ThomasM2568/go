package main

import "fmt"

func main() {
	var notes map[string]int

	notes2 := map[string]int{"a": 10, "b": 20}
	fmt.Println(notes2)
	fmt.Println(notes)
	notes2["c"] = 15
	notes2["a"] = 11
	fmt.Println(notes2)
	fmt.Println(notes2["a"])

	for i := range notes2 {
		fmt.Println("----------")
		fmt.Println(i, " : ", notes2[i])

	}
	delete(notes2, "a")
	fmt.Println(notes2)
}
