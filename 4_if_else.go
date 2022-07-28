package main

import "fmt"

func main() {
	var a int
	a = 5

	if a < 5 {
		fmt.Println("a < 5")
	} else if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a = 5")
	}
}
