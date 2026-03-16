package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "woi, sialan"

	// mengubah menjadi huruf kecil
	fmt.Println("Lowercase :", strings.ToLower(text))

	// mengubah menjadi huruf besar
	fmt.Println("Uppercase :", strings.ToUpper(text))

	// mengecek apkah sting dimulai dengan halo
	fmt.Println("Start with hallo?", strings.HasPrefix(text, "Halo"))

	// mengecek apakah mengandung kata dunia
	fmt.Println("contains dunia", strings.Contains(text, "dunia"))

	// memisahkan string berdasarkan delimiter
	parts := strings.Split("ayam,dog,cat", ",")
	fmt.Println("Split :", parts)

	// mengganti bagian string
	newText := strings.ReplaceAll(text, "sialan", "Golang")
	fmt.Println("Replace:", newText)

}
