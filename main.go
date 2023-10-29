package main

import (
	"fmt"
	"reflect"
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

func Position(arr interface{}, target interface{}) int {
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		v := array.Index(i)
		if v.Interface() == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Hello World")

}
