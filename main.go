package main

import (
	"errors"
	"fmt"
)

func SomeSum(num1 int , num2 int) (int , error) {
	num3 := num1 - num2 
	if num3 < 0 {
		return 0 , errors.New("number 1 must be greater than num2 ")
	}
	return num3 , nil
}

func main() {
	answer , _ := SomeSum(5,2)
	fmt.Printf("Call func and her answer with 5 and 2 is: %d\n" , answer)
}	