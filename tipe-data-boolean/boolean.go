package main

import "fmt"

func main() {
	// deklarasi variabel
	var isRaining bool = true
	var isSunny bool = false
	isLoggedIn := true
	hasPermission := false
	// tampil
	fmt.Println("Apakah Hujan?", isRaining)
	fmt.Println("Apakah Cerah?", isSunny)
	fmt.Println("Apakah Login?", isLoggedIn)
	fmt.Println("Apakah punya akses?", hasPermission)
}
