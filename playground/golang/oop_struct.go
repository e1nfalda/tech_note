package main

import (
	"fmt"
)

type Person struct {
	name, sex string
}

type Policeman struct {
	Person
	station string
}

func NewPoliceman(name, sex, station string) *Policeman {
	return &Policeman{Person{name, sex}, station}
}

func main() {
	police := Policeman{Person{"wang", "male"}, "abc"}
	police2 := &Policeman{Person{"wang", "male"}, "abc"}
	police3 := NewPoliceman("zhang", "femail", "Peak")

	fmt.Println("Hello, playground", police.name, police2.name, police3.name, police3)
}
