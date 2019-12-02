package day2

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func Part1() int {
	var ADD, MUL, END int = 1, 2, 99
	file, err := os.Open("./day2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	strs, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	var values = []int{}
	// Convert from string to int
	for _, v := range strs {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, v)
	}

	values[1] = 12
	values[2] = 2
	for index := 0; index < len(values); index += 4 {
		opCode := values[index]
		firstPosition := values[index+1]
		secondPosition := values[index+2]
		thirdPosition := values[index+3]
		if firstPosition >= len(values) || secondPosition >= len(values) || thirdPosition >= len(values) {
			return values[0]
		}
		if opCode == ADD {
			resulting := values[firstPosition] + values[secondPosition]
			values[thirdPosition] = resulting
		} else if opCode == MUL {
			resulting := values[firstPosition] * values[secondPosition]
			values[thirdPosition] = resulting
		} else if opCode == END {
			return values[0]
		} else {
			log.Fatal("SOMETHING WENT WRONG")
		}
	}
	return values[0]
}
