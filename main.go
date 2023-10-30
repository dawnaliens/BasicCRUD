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

func Unique(arr []int) []int {
	arrLen := len(arr) - 1
	for ; arrLen > 0; arrLen-- {
		for i := arrLen - 1; i >= 0; i-- {
			if arr[arrLen] == arr[i] {
				arr = append(arr[:i], arr[i+1:]...)
				break
			}
		}
	}
	return arr
}

func main() {
	fmt.Println("Hello World")

}
