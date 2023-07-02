package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

var (
	csvFileFlag = flag.String("csv", "problems.csv", "The path to CSV file that contains the problems")
)

func getQuestions(fileName string) []Question {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Unable to open file "+fileName, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, csvErr := csvReader.ReadAll()

	if csvErr != nil {
		log.Fatal("Unable to parse CSV file", csvErr)
	}

	var questions []Question

	for _, value := range records {
		questions = append(questions, Question{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		})
	}

	return questions
}

func main() {
	flag.Parse()

	questions := getQuestions(*csvFileFlag)

	var ans string
	marks := 0

	fmt.Println("Hey There, Let's see how smart you are 👨🏽‍🏫")
	for i, q := range questions {
		fmt.Printf("%d. %s = ", i+1, q.question)
		fmt.Scanln(&ans)

		if ans == q.answer {
			marks += 1
		}
	}

	percent := marks / len(questions) * 100

	fmt.Println("--------------------------------------------------")
	fmt.Println("Your scores are : ", marks, "/", len(questions))
	if percent >= 80 {
		fmt.Println("🥳🥳🥳🥳🥳🔥🔥🔥🔥🔥🎉🎉🎉🎉🎉")
	}
	fmt.Println("--------------------------------------------------")
}
