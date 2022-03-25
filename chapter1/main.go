package main

import "fmt"

func main() {
	fmt.Println("this is from main calling")
	foo()
}

func foo() {
	fmt.Println("this is calling from foo function")
	twoTable()
}
func twoTable() {
	for i := 1; i < 20; i++ {
		fmt.Println("2 *", i, " =", i*2)
	}
	value := 1060
	switch value {
	case 100:
		fmt.Println("The value is 100")
	case 200:
		fmt.Println("The value is morethan 100")
	default:
		fmt.Println("No value is matched")
	}
}
