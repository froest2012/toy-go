package main

import "fmt"

func main() {
    var a,b,c int = 4,89,2
    fmt.Printf("%d %d %d\n", a,b,c)
    d, e := 2, 3
    fmt.Println(d, e)
    fmt.Println("Hello world")
    const LENGTH, WIDTH = 20, 10
    fmt.Println(LENGTH * WIDTH)
    const (
        Unknown = 0
        Female = 1
        Male = 2
    )
    fmt.Println(Unknown)
    fmt.Println(Female)
    const (
        x1 = iota
	x2
	x3
	x4 = "ha"
	x5
	x6 = 100
	x7
	x8 = iota
	x9
    )
    fmt.Println(x1,x2,x3,x4,x5,x6,x7,x8,x9)
}
