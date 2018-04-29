package main

import "fmt"

func main() {

	// declaration
	var msg string
	// assignment
	msg = "Hello Friend"
	// use the variable
	fmt.Println(msg)

	// initialization
	var age int = 27
	// or var age = 27 type is automatically inferred
	// use the variable
	fmt.Print("Your age is : ")
	fmt.Println(age)

	// Shorthand method
	i := 137
	f := 7.14
	s := "Golang is awsome"
	b := true

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", b, b)
}
