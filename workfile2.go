package main

import (
	"fmt"
)

func fyziSe(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b

	if a != nil {
		fmt.Println(*a)
		fmt.Println(*b)
	}

}

func main() {

	var a = 10
	var b = 20

	fmt.Println("Первоначальные данные", a, b)
	fmt.Println("После _________________________________--")
	fyziSe(&a, &b)

}
