package main

import "fmt"

func main() {

	var hari string
	fmt.Println("Silahkan Masukan Hari")
	fmt.Scanln(&hari)

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")

	default:
		fmt.Println("hari tidak diketahui")

	}
}
