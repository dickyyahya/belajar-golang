package main

import "fmt"

func main() {
	var angka1, angka2 int

	// input dari pengguna

	fmt.Print("Masukan angka pertama :")
	fmt.Scanln(&angka1)
	fmt.Print("Masukan angka kedua :")
	fmt.Scanln(&angka2)

	// Latihan perbandingan

	fmt.Println("\n== Hasil Perbandingan ===")
	fmt.Printf("%d==%d %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d!=%d %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d>%d %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d<%d %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d>=%d %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("%d<=%d %v\n", angka1, angka2, angka1 <= angka2)
	fmt.Println("AND:", angka1 > 0 && angka2 > 0)
	fmt.Println("OR :", angka1 > 0 || angka2 > 0)
	fmt.Println("NOT :", !(angka1 > 0))                  // false
	fmt.Println("NOT OR :", !(angka1 > 0 || angka2 > 0)) // false

}
