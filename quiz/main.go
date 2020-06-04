package main

import (
	"fmt"
	"flag"
	"encoding/csv"
	"os"
	"io"
)
type node struct {
	question, answer string
}
func main(){
	filenameptr := flag.String("file", "problems.csv", "name of the csv file having problems")
	flag.Parse()
	filename := *filenameptr

	var a[101] node
	a[100] = node{question:"5+5", answer:"10"}
	var cnt int = 0
	fp, _ := os.Open(filename)
	csvr := csv.NewReader(fp)

	for {
		word, err := csvr.Read()
		if err == io.EOF {
			break
		}
		cnt = cnt+1
		a[cnt] = node{question: word[0], answer: word[1]}
	}
	var i, correct, wrong int
	correct = 0
	wrong = 0;
	var input string
	for i = 1; i <=cnt; i++ {
		fmt.Println(a[i].question)
		fmt.Scanln(&input)
		if input==a[i].answer {
			correct = correct +1
		} else {
			wrong = wrong + 1
		}
	}
	fmt.Println("Number of correct answers:", correct)
	fmt.Println("Number of wrong answers", wrong)
}