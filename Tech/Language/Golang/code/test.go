/* test.go
 */

package main

import "fmt"
import (
	"errors"
	"reflect"
	"time"
)

func testPanic() (result int, err error) {
	//panic("the panic")

	return
}

type mineType struct {
	name string
	id   int
}

func tf(a *mineType) {
	fmt.Println(a.name)
}

func main() {

	b := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println("len", len(b), len(b[len(b)-1]))
	/*
	   var c *[][]int
	   c = &b
	   fmt.Println(b)
	*/
	errors.New("Error")

	var c []int
	fmt.Printf("before addr %p, content: %s\n", &c, c)
	c = append(c, []int{0, 0, 0, 0, 0}...)
	fmt.Printf("after addr %p, content: %s", &c, c)
	copy(c, []int{1, 2, 3, 4, 5})
	fmt.Println("ccccc1: ", c)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("...", reflect.TypeOf(r))
		}
	}()

	/*
	   for x := 3; x > -1; x -- {
	       _ = 10 / x
	   }
	*/

	mt := mineType{name: "helloworld"}

	tf(&mt)

	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}

	fmt.Println("uuuuuuuuuu", []int{1, 2, 3})

	const (
		timeFormat = "2016-01-02 15:04:05"
	)

	td1, _ := time.Parse("2006-01-02", "2018-03-04")
	td2, _ := time.Parse("2006-01-02", "2018-03-05")
	fmt.Println("==========", td1, td2)

	var lista []interface{}
	listb := make([]interface{}, 0)
	fmt.Println(lista, listb)

	my_type := mineType{name: "helloworld", id: 123}
	mstring := fmt.Sprintf("=====%s========", my_type)
	fmt.Println(mstring)
	s1 := fmt.Sprint("ex_keys_%d", "hello", 123)
	fmt.Println(s1)
}
