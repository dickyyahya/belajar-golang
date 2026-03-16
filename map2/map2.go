package main

import "fmt"

func main() {
	mahasiswa := map[string]int{
		"yahya": 80,
		"budi":  70,
		"asep":  60,
	}

	fmt.Println("Nama :", mahasiswa["yahya"])
}
