package main

import (
	"fmt"
  "math"
)

func main() {
	  var f float32 = 16777216
	  fmt.Println(f == f + 1)
    
    var g float64 = 1
    for i := 0;  g != g + 1; i++ {
        fmt.Println(i)
    }

    const e = 2.71828
    const Avogadro = 6.02214129e23
    const Planck = 6.2606957e-34

    for x := 0; x < 8; x++ {
        fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
    }

    var z float64
    fmt.Println(z, -z, 1/z, -1/z, z/z)

    nan := math.NaN()
    fmt.Println(nan == nan, nan < nan, nan > nan)
}
