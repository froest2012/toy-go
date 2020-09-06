package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println(a)
	if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}
	bb := [4]string{"a", "b", "c", "d"}
	for i, k := range bb {
		fmt.Println(i, k)
	}
}
