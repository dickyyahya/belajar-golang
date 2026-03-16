package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:] //mengambil seluruh element

	fmt.Println("Ini adalah s1 :", s1)

	s2 := arr[:3]
	fmt.Println("ini adalah s2 :", s2)

	s3 := arr[2:]
	fmt.Println("ini adalah S3", s3)

	s4 := arr[1:4]
	fmt.Println("ini adalah s4 :", s4)

}
