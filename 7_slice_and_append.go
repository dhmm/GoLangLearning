package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr1)
	fmt.Println(arr2)

	//Slices are parts of the array that point on the arrays
	//If we change anyhthing also original array will change
	nArr1 := arr1[0:5]
	nArr2 := arr2[5:10]

	fmt.Println("Result : ", "\n", nArr1, "\n", nArr2)

	nArr1[0] = 99

	fmt.Println("Result after change on slice\n", arr1, "\n", nArr1)

	//Print type of varibles
	fmt.Println(reflect.ValueOf(nArr1).Kind())
	fmt.Println(reflect.ValueOf(arr1).Kind())

}
