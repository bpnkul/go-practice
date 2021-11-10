package main

import "fmt"

func main() {
	intArray := []int{0,1,2,3,4,5,6,7,8,9}

	for i := range intArray {
		if i%2 == 0{
			fmt.Println("Number is even: ", i)
		}else{
			fmt.Println("Number is odd: ", i)
		}
	}
}