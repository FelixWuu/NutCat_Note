package main

import (
	"fmt"
	"reflect"
)

func SliceDeepEqualDemo() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{3, 2, 1}

	// 比较 a 和 b 是否相等
	if reflect.DeepEqual(a, b) {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	// 比较 a 和 c 是否相等
	if reflect.DeepEqual(a, c) {
		fmt.Println("a and c are equal")
	} else {
		fmt.Println("a and c are not equal")
	}
}
