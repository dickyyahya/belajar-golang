package main

import "fmt"

func main() {
	// x := 10
	var angka int
	fmt.Print("Masukan angka :")
	fmt.Scanln(&angka)

	if angka >= 90 {
		fmt.Println(angka, "GG")
	} else if angka >= 70 {
		fmt.Println(angka, "Baik")
	} else {
		fmt.Println(angka, "Gagal")
	}
}
