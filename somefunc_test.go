package main

import (
	"fmt"
	"testing"
)

func TestFunc(test *testing.T) {
	expacted := 3
	answer , _ := SomeSum(5,2)
	if expacted != answer {
		fmt.Print("Fall test")
		test.Error("expacted and answer are different")
	}
}