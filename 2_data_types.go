package main

import "fmt"

func main() {
	//declaring integer
	var a int
	a = 5
	//printing integer
	fmt.Println(a)
	//printing with string
	fmt.Println("a =", a, "\t\ta + 2 =", (a + 2), "\ta ^ 2 =", a^2)

	//32 bit float
	var f32 float32
	f32 = 5.632456
	fmt.Println(f32)

	//32 bit float
	var f64 float64
	f64 = 5.6324564567890765432425678
	fmt.Println(f64)

	//strings
	var name string
	name = "Moutlou"
	fmt.Println(name)
	var surname string
	surname = "NCM"
	fmt.Println(surname)

	//string concatenation
	var snn string
	snn = name + " , " + surname
	fmt.Println(snn)

	//CONSTANT
	const b = 5
	fmt.Println(b)

	//mix var declaration
	var v1, v2, v3 = 5, "name", 6.78
	fmt.Println(v1, " , ", v2, " , ", v3)

	//DECLARATION AND ASSIGNMENT
	declareAndAssign := 5
	fmt.Println(declareAndAssign)

}
