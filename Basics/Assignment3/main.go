package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open("temp.txt")
	if(error != nil){
		fmt.Println("Error: ", error);
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)

}