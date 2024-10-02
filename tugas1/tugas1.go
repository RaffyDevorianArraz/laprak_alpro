package main

import "fmt"

func main() {
	var (satu, dua, tiga, temp string)
	
	fmt.Print("Nama : ")
	fmt.Scanln(&satu)
	fmt.Print("Nim: ")
	fmt.Scanln(&dua)
	fmt.Print("Prodi: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
