package main

import "fmt"

func main() {
	fmt.Print("Enter First number: ") //Print function is used to display output in same line
	var num1 int
	fmt.Scanln(&num1) // Take input from user
	fmt.Print("Enter Second number: ")
	var num2 int
	fmt.Scanln(&num2)
	var sum = num1 + num2
	fmt.Println(sum) // Addition of two number
}
