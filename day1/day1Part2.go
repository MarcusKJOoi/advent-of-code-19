package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calculateFuel(fuel int) int {
	currentFuel := int(math.Floor(float64(fuel)/3.0)) - 2
	if currentFuel <= 0 {
		return 0
	}
	return currentFuel + calculateFuel(currentFuel)
}

func Part2() {
	file, err := os.Open("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	scanner.Split(bufio.ScanWords)
	next := scanner.Scan()
	sum := 0
	for next {
		txt := scanner.Text()
		fuel, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		sum += calculateFuel(fuel)
		next = scanner.Scan()
	}
	fmt.Println("Day 2 Part 2", sum)
}
