package main

import "fmt"

func main() {

	// angka := 10
	var angka int
	fmt.Println("masukan angka")
	fmt.Scanln(&angka)

	switch {
	case angka >= 90:
		fmt.Println("A") //>=90

	case angka >= 75 && angka <= 89:
		fmt.Println("B") //>= 75 && <=89

	case angka >= 60 && angka <= 74:
		fmt.Println("C") // >=60 && <=74

	default:
		fmt.Println("Bubar")

	}
}
