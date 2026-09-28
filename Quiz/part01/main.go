package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "problems", "use it to parse the problems file name")
	flag.Parse()

	var input int
	var score int

	type Quiz struct {
		fn  int
		sn  int
		res int
	}

	// opening my csv file
	file, err := os.Open(filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}

	// extracting its data
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// reading through the extracted data
	r := csv.NewReader(strings.NewReader(string(data[:count])))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var quizzes []Quiz

	// splitting the csv fields into num1, num2 and res
	for i, row := range records {

		sum := strings.TrimSpace(row[0])

		parts := strings.Split(sum, "+")
		if len(parts) != 2 {
			log.Fatalf("Error: '%s' field has no parts separated by '+'", sum)
		}

		num1, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Printf("Row %d: Failed to parse sum '%s': %v", i, parts[0], err)
			continue
		}

		num2, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Printf("Row %d: Failed to parse sum '%s': %v", i, parts[1], err)
			continue
		}

		res, err := strconv.Atoi(strings.TrimSpace(row[1]))
		if err != nil {
			log.Printf("Row %d: Failed to parse sum '%s': %v", i, row[1], err)
			continue
		}

		quizzes = append(quizzes, Quiz{
			fn:  num1,
			sn:  num2,
			res: res,
		})
	}

	fmt.Println("-------STARTING QUIZ-------")

	for i, q := range quizzes {
		fmt.Printf("Quiz %d: %d + %d \n", i+1, q.fn, q.sn)

		fmt.Scanf("%d\n", &input)

		if input == q.res {
			fmt.Println("CONGRATS, RIGHT ANSWER")
			score++
		}
	}

	fmt.Printf("QUIZ ENDED, YOUR SCORE: %d \n", score)
}
