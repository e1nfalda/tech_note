package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello,world"

	//fmt.Println("----------", strings.Split(s, ","))

	for _, ss := range strings.Split(s, ",") {
		fmt.Println("ss: ", ss)
	}
}
