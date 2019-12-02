package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.txt")
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
		// fmt.Println(txt)
		i, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		sum += int(math.Floor(float64(i)/3.0)) - 2
		next = scanner.Scan()
	}
	fmt.Println(sum)

}
