package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	isTimeout := false
	var quizFile string
	var timeout, totalCorrect int
	flag.StringVar(&quizFile, "q", "quiz.csv", "Quiz csv file")
	flag.IntVar(&timeout, "t", 30, "timeout")
	flag.Parse()
	records := readQuizFile(quizFile)
	for i, record := range records {
		fmt.Printf("%d. %s = \n", i+1, record[0])
		input := make(chan string, 1)
		go getAnswer(input)
		select {
		case ans := <-input:
			if ans == strings.TrimSpace(record[1]) {
				totalCorrect++
			}
		case <-time.After(time.Duration(timeout) * time.Second):
			isTimeout = true
		}
		if isTimeout {
			fmt.Println("You didn't answer in time")
			break
		}
	}
	fmt.Printf("Your quiz score is %d/%d", totalCorrect, len(records))
}

func readQuizFile(quizFile string) [][]string {
	file, err := os.Open(quizFile)
	if err != nil {
		fmt.Println("Error while open quiz file: ", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error while reading quiz file: ", err.Error())
	}
	return records
}

func getAnswer(input chan string) {
	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Can't get answer ", err.Error())
	}
	input <- strings.TrimSpace(ans)
}
