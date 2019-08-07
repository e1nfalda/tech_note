/* helloWorld.go
 *
 */
package main

import "fmt"

func main() {
	a := 10
	var pa *int = &a
	array := [5]int{1, 2, 3}
	var parray *[5]int
	parray = &array

	fmt.Println("hello world!", a)
	fmt.Println("var: ", a, pa, *pa)
	fmt.Println("array", array, *parray, parray, &array)
}
