package main

import "fmt"

func main() {
	fmt.Print("start\n")
	testFun1()

	fmt.Print("end")
}

func testFun1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("-----------recover:", r, "\n")
		}
	}()
	testFun()
}

func testFun() {
	defer fmt.Println("aaaaaaaaaaaa\n")
	panic("this is a panic")
}
