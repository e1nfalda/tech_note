package main

import "fmt"

type myType struct {
	str string
}

func (p myType) String() string {
	return fmt.Sprintf("%v-%T", p.str, p.str)
}

func main() {
	p := myType{"abc"}
	fmt.Println(p)
}
