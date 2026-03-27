package main

import "fmt"

func main() {
	const USDinUER float64 = 0.86
	const USDinRUB float64 = 82.13
	EURinRUB := 1 * (USDinUER / USDinRUB)
	fmt.Println(EURinRUB)
}
