package main

import (
	"encoding/csv"
	"os"
)

type quizQuestion struct {
	question string
	answer   string
}

func check_for_errors(err error) {
	if err != nil {
		panic(err)
	}
}

func parse_data(data [][]string) []quizQuestion {
	var questions []quizQuestion

	for _, row := range data {
		questions = append(
			questions,
			quizQuestion{question: row[0], answer: row[1]},
		)
	}

	return questions
}

func read_csv(file_path string) {
	file, err := os.Open(file_path)

	if err != nil {
		check_for_errors(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	data, err := reader.ReadAll()
	if err != nil {
		check_for_errors(err)
	}

}

func main() {
	read_csv("/src/app/exercise_1/problems.csv")
}
