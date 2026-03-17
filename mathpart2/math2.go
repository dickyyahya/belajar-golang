package main

import "fmt"

func main() {
	a := 5
	a = 2

	a += 2 //x= a+2
	fmt.Println("a+=2", a)

	a -= 2 //a= a-2
	fmt.Println("a-=2", a)

	a *= 2 //a= a*2
	fmt.Println("a*=2", a)

	a /= 2 //a= a/2
	fmt.Println("a/=2", a)

	a %= 2 //a= a%2
	fmt.Println("a%=2", a)
}
