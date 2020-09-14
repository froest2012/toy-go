package main

import (
    "fmt"
)

type Currency int

const (
    USD Currency = iota
    EUR
    GBP
    RMB
)

func main() {
    var a [3]int
    fmt.Println(a[0])
    fmt.Println(a[len(a)-1])

    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }
    for _, v := range a {
        fmt.Printf("%d\n", v)
    }
    var q [3]int = [3]int{1,2,3}
    var r [3]int = [3]int{1,2}
    fmt.Println(r[2], q[1])
    x := [...]int{1,2,3,4,5}
    fmt.Printf("%T\n", x)
    fmt.Println(x)
    symbol := [...]string{USD:"$", EUR:"€", GBP: "￡", RMB: "￥"}
    fmt.Println(RMB, symbol[RMB])
}
