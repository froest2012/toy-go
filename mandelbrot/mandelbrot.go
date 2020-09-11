package main

import (
//    "fmt"
    "math/cmplx"
    "image"
    "image/color"
    "os"
    "image/png"
)

func main() {
/*
    var x complex128 = complex(1, 2)
    var y complex128 = complex(3, 4)
    fmt.Println(x*y)
    fmt.Println(real(x*y))
    fmt.Println(imag(x*y))
    fmt.Println(1i*1i)
    fmt.Println(cmplx.Sqrt(-1))
*/
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            //return color.Gray{255 - contrast*n}
            return color.RGBA{uint8(n), uint8(n), 255, 255}
        }
    }
    //return color.Black
    return color.RGBA{uint8(55), uint8(39), 92, 78}
}
