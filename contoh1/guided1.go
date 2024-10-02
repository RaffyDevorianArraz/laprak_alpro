package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	sum := a + b + c + d + e
	fmt.Printf("Hasil penjumlahan %d + %d + %d + %d + %d adalah %d\n", a, b, c, d, e, sum)
}
