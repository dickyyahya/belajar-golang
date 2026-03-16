package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"nama":  "yahya",
		"kelas": "B",
	}
	fmt.Println("Nama :", mahasiswa["nama"])
	fmt.Println("Kelas :", mahasiswa["kelas"])

	mahasiswa["jurusan"] = "Informatika"
	fmt.Println("Jurusan :", mahasiswa["jurusan"])
}
