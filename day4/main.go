package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	return score
}

func DayFourStar1() {

	lines := importDataByFile()
	sum := 0
	for _, line := range lines {

		tab := strings.FieldsFunc(line, func(r rune) bool { return r == ':' || r == '|' })
		left := strings.FieldsFunc(tab[1], func(r rune) bool { return r == ' ' })
		right := strings.FieldsFunc(tab[2], func(r rune) bool { return r == ' ' })
		match := MatchNumber(left, right)
		sum = sum + CalculPoints(match)

	}

	fmt.Println(sum)

}

type Card struct {
	Number      int
	nb_passages int
}

func DayFourStar2() {

	lines := importDataByFile()
	var additionnal []Card
	for i := 1; i < len(lines)+1; i++ {
		additionnal = append(additionnal, Card{Number: i, nb_passages: 1})
	}

	for _, line := range lines {

		tab := strings.FieldsFunc(line, func(r rune) bool { return r == ':' || r == '|' })
		card_n, err := strconv.Atoi(strings.FieldsFunc(tab[0], func(r rune) bool { return r == ' ' })[1])
		if err != nil {

			log.Fatal(err)

		}
		for i := 0; i < additionnal[card_n-1].nb_passages; i++ {
			left := strings.FieldsFunc(tab[1], func(r rune) bool { return r == ' ' })
			right := strings.FieldsFunc(tab[2], func(r rune) bool { return r == ' ' })
			match := MatchNumber(left, right)

			for i := 0; i < len(match); i++ {

				additionnal[card_n+i].nb_passages = additionnal[card_n+i].nb_passages + 1
			}

		}
	}

	sum := 0
	for _, card := range additionnal {
		sum += card.nb_passages
	}
	fmt.Println(sum)

}

func main() {

	startTime := time.Now()
	fmt.Println("Day 4")
	fmt.Println("Star 1")
	DayFourStar1()
	fmt.Println("Star 2")
	DayFourStar2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
