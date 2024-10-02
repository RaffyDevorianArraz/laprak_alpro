package main

import "fmt"

func main() {
    var r, l float64
    var PHI float64 = 3.14 

    fmt.Print("Masukkan jari-jari: ")
    fmt.Scanln(&r)

    l = PHI * r * r

    fmt.Println("Luas lingkaran:", l)
}
