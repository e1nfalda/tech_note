package main

import "fmt"
import "reflect"

func main() {
	slice1 := make([]int, 2)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, slice[:3], reflect.TypeOf(slice), slice1)
	copy(slice1, slice[:3])

	fmt.Println(slice1)
}
