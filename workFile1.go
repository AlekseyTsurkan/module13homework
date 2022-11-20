package main

import "fmt"

func tyPlace(a, b int) {

	if a%2 == 0 || b%2 == 0 {
		fmt.Println("Числа четные,", a, b)
		fmt.Println(a + b)
		fmt.Println(a - b)
	} else {
		fmt.Println("Числа не четные")
	}
}
func main() {
	tyPlace(5, 7)
	tyPlace(2, 2)
	tyPlace(10, 20)

}
