package main

import "fmt"

func main() {
	var x int = 10
	var y = 20
	fmt.Println(x + y)

	var first_name string = "Zachary"
	const (
		last_name = "Cheng"
	)

	fmt.Println(first_name + last_name)
	first_name = "TEst!"

	fmt.Printf("Hello, %s!\n", "world")
}
