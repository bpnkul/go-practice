package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := getRandomNumbersToSort(10)
	printNumbers(numbers)
	bubbleSort(numbers)
	printNumbers(numbers)
}
func getRandomNumbersToSort(size int) [] int{
	var numbers = []int{}
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<size; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	return numbers
}

func printNumbers(numbers []int ){
		fmt.Println(numbers)
}

func bubbleSort(numbers []int){
	for i:=0; i<len(numbers); i++{
		for j:=i; j<len(numbers); j++{
			if numbers[i] > numbers[j]{
				numbers[i], numbers[j] = numbers[j], numbers[i]
				printNumbers(numbers)
			}
		}
	}
}