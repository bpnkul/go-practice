package main

import "fmt"

func main(){
	
	// colors := map[string]string{
	// 	"purple":"#PURPLE",
	// 	"PINK": "#PINK",
	// }
	//var colors map[string]string
	colors := make(map[string]string)
	
	colors["red"] = "#RED"
	colors["blue"] = "#BLUE"
	printMap(colors)
	colors["white"] = "#FFFFF"
	printMap(colors)
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string){
	fmt.Println("------------")
	for color, hex := range c {
		fmt.Println(color," : " ,hex)
	}
}