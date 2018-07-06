package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var (
	flagFilePath string
)

type question struct{ question, answer string }

func init() {
	flag.StringVar(&flagFilePath, "file", "questions.csv", "/path/to/file.csv")
	flag.Parse()
}

func main() {
	questions, err := readCSV(flagFilePath)
	if err != nil {
		fmt.Printf("Cannot parse file %s: %v\n", flagFilePath, err)
	}

	fmt.Printf("read questions: %v", questions)
}

func readCSV(filePath string) ([]question, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot open: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = 2
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read problem: %v", err)
	}

	out := make([]question, len(lines))
	for lineNum, line := range lines {
		out[lineNum] = question{line[0], line[1]}
	}

	return out, nil
}
