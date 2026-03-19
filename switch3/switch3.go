package main

import "fmt"

func main() {

	var panjang string
	fmt.Println("Masukan Kata")
	fmt.Scanln(&panjang)

	switch panjang := len(panjang); {
	case panjang > 10:
		fmt.Println("kata terlalu panjang")

	case panjang >= 5:
		fmt.Println("kata terlalu sedang")

	default:
		fmt.Println("kata pendek")
	}
}
