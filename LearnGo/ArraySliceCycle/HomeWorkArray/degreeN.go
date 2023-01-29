// Домашнее задание. Как бы вы реализовали функцию, которая в качестве аргумента принимает число n,
// и выводит в консоль все числа от 2^1 до 2^n ?
// Выполнено.
package main

import (
	"fmt"
	"math"
)

func main() {
	var arrayN []int
	arrayN = degreeN(arrayN)
	for index := range arrayN {
		fmt.Printf("%d. %d\n", index+1, arrayN[index])
	}
}
func degreeN(arrayN []int) []int {
	var num int
	n := 41
	for i := 1; i <= n; i++ {
		num = int(math.Pow(float64(2), float64(i)))
		arrayN = append(arrayN, num)
	}
	return arrayN
}
