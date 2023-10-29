package main

import (
	"fmt"
)

func ExistInArray(target int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if target == array[i] {
			return true
		}
	}
	return false
}

func ExistInAnArray(target int, array []int) bool {
	for _, v := range array {
		if target == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Hello World")
	
}
