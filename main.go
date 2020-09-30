package main

import "fmt"

func main() {
	var numberE, acumulator float32
	var limit int
	fmt.Scan(&limit)
	numberE = 1
	acumulator = 1
	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			acumulator = acumulator * float32(j)
		}
		numberE = numberE + (1 / acumulator)
		acumulator = 1
	}
	fmt.Print(numberE)
}
