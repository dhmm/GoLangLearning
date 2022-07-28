package main

import "fmt"

func main() {
	var a = 10
	switch a {
	case 1:
		fmt.Println("a = 1")
		break
	case 2:
		fmt.Print("a = 2")
		break
	default:
		fmt.Println("a <> 1 && a <> 2")
		break

	}
}
