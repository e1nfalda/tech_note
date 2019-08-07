package main

import (
	"fmt"
	"runtime"
	// "time"
)

func main() {
	names := []string{"lily", "yoyo", "cersei", "rose", "annei"}
	for _, name := range names {
		go func(name string) {
			fmt.Println(name)
		}(name)
	}
	runtime.GOMAXPROCS(1)
	runtime.Gosched()
}
