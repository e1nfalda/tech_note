/* json.go
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	//jsonBlob = []byte(test_str)
	// test_str := `[{"position":763, "test": 123},{"position":587, "test": 123},{"position":610, "test": 123}]`
	test_str := `[{"position":763}]`
	type Parameter struct {
		position int `tag:"position"`
	}

	var ads []Parameter
	err := json.Unmarshal([]byte(test_str), &ads)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(ads)
	t1 := Parameter{11}
	fmt.Println(t1.position)

	js, err := simplejson.NewJson([]byte(`{"t":"1563897600"}`))
	fmt.Println("MustInt", js.Get("t").MustInt())

	/*
		var animals []Animal
		err := json.Unmarshal(jsonBlob, &animals)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", animals)
	*/
}
