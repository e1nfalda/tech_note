package main

import (
	"fmt"
	"reflect"
)

func pointer_func(a *[]int) {
	fmt.Println("pointer func:", *a, a)
}

func normal_func(a [5]int) {
	fmt.Println("normal func: ", a, &a)
}

func main() {
	// var s [:]int
	var s = make([]int, 5)
	//var s = [5]int{}

	fmt.Println(reflect.TypeOf(s))

	s[1] = 1

	// pointer_func(&s)
	//normal_func(s)
	//s[2] = 2
	// pointer_func(&s)
	//normal_func(s)
	for index, v := range s {
		fmt.Println(index, ": ", v)
	}

	m := make(map[string]int)
	m["first"] = 1
	fmt.Println(m)

	var v = "123"
	fmt.Println(int(v))
}
