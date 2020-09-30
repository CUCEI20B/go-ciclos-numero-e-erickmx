package main

import "fmt"

func main() {
	var numberE, acumulator float64
	numberE = 1
	acumulator = 1
	for i := 1; i < 1000000; i++ {
		for j := 2; j <= i; j++ {
			acumulator = acumulator * float64(j)
		}
		numberE = numberE + (1 / acumulator)
		acumulator = 1
	}
	fmt.Print(numberE)
}
