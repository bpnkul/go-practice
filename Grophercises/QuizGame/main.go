package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func main(){

	puzFile := flag.String("input","problems.csv", "Quiz input file.")
	flag.Parse()
	fmt.Println("Quiz File : ", *puzFile)

	f, err := os.Open(*puzFile)
	if err != nil {
		exit(fmt.Sprintf("Error while opening file : %v\n", err))
	}
	fmt.Println("File opened.")
	defer f.Close()

	rd := csv.NewReader(f)
	rio := bufio.NewReader(os.Stdin)
	score := 0
	total := 0
	for {		
		line, err := rd.Read()
		if err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println("Error while reading record :", err)
			continue
		}
		fmt.Printf("Question : %v : Response :", line[0])
		textRead,_ := rio.ReadString('\n')
		text := strings.TrimRight(textRead,"\n") 
		text = strings.TrimRight(text,"\r")
		total++
		if strings.EqualFold(text,line[1]) {
			fmt.Printf("Congratulations. Correct answer : %v\n", text)
			score++
		}else{
			fmt.Printf("Sorry. Your answer : %v, Correct answer : %v\n", text, line[1])
		}
	}
	fmt.Printf("Your score : %v / %v", score, total)
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func parseProble(line []string) (q string, a string){
	fmt.Println("Quetion length : ", len(line))
	if(len(line) != 2){
		q := nil
		a := nil
		return *q, *a
	}
	return line[0], line[1]
}