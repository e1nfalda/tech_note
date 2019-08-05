/* channel.go
 */
package main

import (
    "time"
    "fmt"
)


func main() {
    c := make(chan int)

    go func() {
        for i:=0; i<3; i++ {
            time.Sleep(1 * time.Second)
            // fmt.Println("i: ", i)
            c <- i
        }
    }()

    for {
        v := <-c
        fmt.Println("=====", v)
    }
}
