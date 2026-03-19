package main

import "fmt"

func main() {
	//loop biasa
	//loop dari 1 sampai 5
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Println("Angka ke - ", i)
	// 	}
	// }
	//infiniti loop
	// for {
	// 	fmt.Println("muter terus kunyuk")

	// }

	// loop dengan range
	// buah := []string{"melon", "jambu", "gedang"}
	// for index, item := range buah {
	// 	fmt.Printf("index %d = %s\n", index, item)
	// }

	// for dengan break dan continue
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue //skip angka 3
	// 	}
	// 	if i == 5 {
	// 		break //berhenti di angka 5
	// 	}
	// 	fmt.Println(i)
	// }

	// 1. for standar untuk mencetak angka 1 ->5
	fmt.Println("\n1. for standar (i = 1; i<=5;i++)")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d", i)
	}

	//2. while-style mencetak angka sampai 10
	fmt.Println("\n2. while-style (hanya kondisi):")
	j := 2
	for j <= 10 {
		fmt.Printf("%d", j)
		j += 2 //j=j+2
	}
	fmt.Println()

	// 3. infinite loop dengan break menghitung mundur 5
	fmt.Println("\n3. infinite loop dengan break:")
	k := 5
	for {
		fmt.Printf("%d", k)
		k--
		if k == 0 {
			break //keluar dari loop
		}
	}
	fmt.Println()

	// 4. Range loop iterasi elemen slice
	fmt.Println("\n4, range loop untuk slice:")
	buah := []string{"nanas", "apel", "melon"}
	for index, value := range buah {
		fmt.Printf("index %d: %s\n", index, value)
	}
}
