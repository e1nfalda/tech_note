package main

import "fmt"

func main() {

	for {
		var a = []int{1, 2, 3, 4, 5, 6, 7}
		fmt.Println("", a)
		a = a[6:]
		fmt.Println("", a)
	}
}
