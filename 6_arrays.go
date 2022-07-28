package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println(arr)
	fmt.Println(arr[0])

	//declaration and assignment
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//declaration without size and assignment
	arr3 := [...]int{20, 40, 60, 80, 100}
	fmt.Println(arr3)

	var arr4 = [...]int{25, 45, 65, 85, 105}
	fmt.Println(arr4)

	//LENGTH OF THE ARRAY
	length := len(arr)
	fmt.Println(length)

}
