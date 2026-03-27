package main

import "fmt"

func main() {
	const USDinEUR float64 = 0.86
	const USDinRUB float64 = 82.13
	EURinRUB := USDinRUB / USDinEUR
	fmt.Println(EURinRUB)
}
