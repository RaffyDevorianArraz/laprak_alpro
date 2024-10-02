package main

import "fmt"

func main() {
	var x float64
	fmt.Scanln(&x)
	hasil := 2/(x+5) + 5
	fmt.Println(hasil)
}
