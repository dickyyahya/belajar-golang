package main

import "fmt"

func main() {

	number := [5]int{10, 20, 30, 40, 50}
	fmt.Println("jumlah elemen :", len(number))
	fmt.Println("index ke-1 :", number[1])
	number[1] = 70
	fmt.Println("index ke-1:", number[1])

	fmt.Println("Ini adalah array :")

	for index, value := range number {
		fmt.Println("isi index ke :", index, "=", value)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	fmt.Println("arr1 == arr2 :", arr1 == arr2)
	fmt.Println("arr 1 != arr2 :", arr1 != arr2)
}
