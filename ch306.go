package main

import (
    "fmt"
    "time"
)

const (
    e  = 2.71828182845904523536028747135266249775724709369995957496696763
    pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

const IPv4Len = 4
const noDelay time.Duration = 0
const timeout = 5 * time.Minute

const (
    a = 1
    b
    c = 2
    d
)

type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

type Flags uint

const (
    FlagUp Flags = 1 << iota
    FlagBroadcast
    FlagLoopback
    FlagPointToPoint
    FlagMulticast
)

const (
    _= 1 << (10 * iota)
    KiB
    MiB
    GiB
    PiB
    EiB
    ZiB
    YiB
)

func main() {
    fmt.Printf("%T %[1]v\n", noDelay)
    fmt.Printf("%T %[1]v\n", timeout)
    fmt.Printf("%T %[1]v\n", time.Minute)
    fmt.Println(a, b, c, d)
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
    fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
    var v Flags = FlagMulticast | FlagUp
    fmt.Printf("%b %t\n", v, IsUp(v))
    TurnDown(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    SetBroadcast(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    fmt.Printf("%b %t\n", v, IsCast(v))
}
/*
func parseIPv4(s string) IP {
    var p [IPv4Len]byte
}
*/

func IsUp(v Flags) bool {
    return v & FlagUp == FlagUp
}

func TurnDown(v *Flags) {
    *v &^= FlagUp
}

func SetBroadcast(v *Flags) {
    *v |= FlagBroadcast
}

func IsCast(v Flags) bool {
    return v & (FlagBroadcast | FlagMulticast) != 0
}


