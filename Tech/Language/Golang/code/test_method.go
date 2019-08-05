/* test_method.go
 */

package main

import (
	"fmt"
)

type Square struct {
	width  int
	height int
}

func (square Square) area() int {
	return square.width * square.height
}

func (square Square) pwidth() {
	fmt.Println(square.width)
}

func main() {
	//square := Square{width:10, height:5} // 另一种初始化方式
	square := Square{10, 10}

	fmt.Println(square.area())
	square.pwidth()
}
