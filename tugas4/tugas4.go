package main

import "fmt"

func main() {
    var fahrenheit float64

    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&fahrenheit)

    celsius := (fahrenheit - 32) * 5.0 / 9.0 

    fmt.Printf("Suhu dalam Celsius adalah %.2f\n", celsius)
}
