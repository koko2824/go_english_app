package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)


// ベタ書き
func main() {
	file, err := os.Open("./sheet01.csv")
	if err != nil {
		panic(err)
	}
	file.Close()

	rand.Seed(time.Now().UnixNano())

	reader := csv.NewReader(file)
	var lines [][]string
	score := 0

	for {
		line, err := reader.Read()
		lines = append(lines, line)
		if err != nil {
			break
		}
	}
	var incorrect [][]string

	for i := 0; i < 10; i++ {
		var ans string
		line := lines[rand.Intn(len(lines) -1)]
		fmt.Println("----次の単語の意味を答えなさい----")
		fmt.Println(line[1]+":")
		_, _ = fmt.Scan(&ans)

		if ans == line[0] {
			score ++
		} else {
			incorrect = append(incorrect, line)
		}
	}

	fmt.Println("正解数:", score)
	fmt.Println("--間違えた単語--")

	for _, v := range incorrect {
		fmt.Println(v[0], ":", v[1])
	}
	fmt.Println(lines)
}