package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 10
	var b float64 = float64(a) //ubah dari int ke float64

	var score int = 95
	var text string = strconv.Itoa(score) //int to string

	fmt.Println("Nilai a :", a)
	fmt.Println("Nilai b:", b)

	fmt.Println("Nilai ujian :", text)

}
