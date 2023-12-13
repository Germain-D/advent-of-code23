package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func importDataByFile() []string {
	fmt.Println("Importing data...")
	/* create tab with rows*/
	var tab []string
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	tab = strings.Split(string(content), "\n")
	fmt.Println("Done.")
	return tab
}

func MatchNumber(left []string, right []string) []string {
	var match []string
	for _, l := range left {
		for _, r := range right {
			if l == r {
				match = append(match, l)
			}
		}
	}
	return match
}
func contains(elems []int64, v int64) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
func CalculPoints(match []string) int {
	score := 0
	for i := 0; i < len(match); i++ {
		if i == 0 {
			score += 1
		} else {
			score = score * 2
		}
	}

	fmt.Println(score)
	return score
}

func DayFourStar1() {

	lines := importDataByFile()
	sum := 0
	for _, line := range lines {
		fmt.Println(line)
		tab := strings.FieldsFunc(line, func(r rune) bool { return r == ':' || r == '|' })
		left := strings.FieldsFunc(tab[1], func(r rune) bool { return r == ' ' })
		right := strings.FieldsFunc(tab[2], func(r rune) bool { return r == ' ' })
		match := MatchNumber(left, right)
		sum = sum + CalculPoints(match)

	}

	fmt.Println(sum)

}

func main() {

	startTime := time.Now()
	DayFourStar1()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
