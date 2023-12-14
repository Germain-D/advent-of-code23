package main

import (
	"fmt"
	"log"
	"math"
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

func RockNRoll(column string) string {

	if strings.Contains(column, "O") {
		/*slit # */
		var new_cols []string
		temp_cols := strings.Split(column, "#")
		for i := 0; i < len(temp_cols); i++ {
			/* get all index of O */
			var new_col string
			o := strings.Count(temp_cols[i], "O")
			for j := 0; j < len(temp_cols[i]); j++ {
				if j < o {
					new_col = new_col + "O"
				} else {
					new_col = new_col + "."
				}
			}
			new_cols = append(new_cols, new_col)
		}

		return strings.Join(new_cols, "#")
	}
	return column
}

func Day14Star1() {

	lines := importDataByFile()
	/* create tab with columns*/
	var tab []string
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " => ")...)
	}
	for i := 0; i < len(tab); i++ {
		tab[i] = ""
		for j := 0; j < len(lines); j++ {

			tab[i] = tab[i] + string(lines[j][i])
		}
	}
	sum := 0
	var rockNRoll []string
	for _, col := range tab {
		rockNRoll = append(rockNRoll, RockNRoll(col))

	}
	for j := 0; j < len(rockNRoll); j++ {
		for i := 0; i < len(rockNRoll[j]); i++ {
			if string(rockNRoll[j][i]) == "O" {
				sum += int(math.Abs(float64(i - len(rockNRoll[j]))))
			}
		}
	}

	fmt.Println(sum)

}

func NorthPole(tab *[]string, lines []string) {
	for i := 0; i < len(*tab); i++ {
		(*tab)[i] = ""
		for j := 0; j < len(lines); j++ {

			(*tab)[i] = (*tab)[i] + string(lines[j][i])
		}
	}

	//fmt.Println("NorthPole", tab)

	var rockNRoll []string
	for _, col := range *tab {
		rockNRoll = append(rockNRoll, RockNRoll(col))
	}

	//fmt.Println("NorthPole", rockNRoll)
	(*tab) = rockNRoll
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SouthPole(tab *[]string) {
	lines := make([]string, len((*tab)))
	copy(lines, (*tab))
	for i := 0; i < len((*tab)); i++ {
		(*tab)[i] = ""
		for j := 0; j < len(lines); j++ {

			(*tab)[i] = (*tab)[i] + string(lines[j][i])
		}
	}
	//fmt.Println("SouthPole", tab)
	/* reverse rows int tab */
	for i := 0; i < len((*tab)); i++ {
		(*tab)[i] = Reverse((*tab)[i])
	}

	//fmt.Println("SouthPole", tab)
	var rockNRoll []string
	for _, col := range *tab {
		rockNRoll = append(rockNRoll, RockNRoll(col))
	}

	//fmt.Println("SouthPole", rockNRoll)
	(*tab) = rockNRoll
}

func WestPole(tab *[]string) {

	lines := make([]string, len((*tab)))
	copy(lines, (*tab))

	for i := 0; i < len((*tab)); i++ {
		(*tab)[i] = ""
		for j := 0; j < len(lines[i]); j++ {

			(*tab)[i] = (*tab)[i] + string(lines[j][i])
		}
	}

	//fmt.Println("WestPole", tab)
	var rockNRoll []string
	for _, col := range *tab {
		rockNRoll = append(rockNRoll, RockNRoll(col))
	}

	//fmt.Println("WestPole", rockNRoll)
	(*tab) = rockNRoll
}

func EastPole(tab *[]string) {
	lines := make([]string, len((*tab)))
	copy(lines, (*tab))

	for i := 0; i < len((*tab)); i++ {
		(*tab)[i] = ""
		for j := 0; j < len(lines[i]); j++ {

			(*tab)[i] = (*tab)[i] + string(lines[j][i])
		}
	}

	/* reverse rows int tab */
	for i := 0; i < len((*tab)); i++ {
		(*tab)[i] = Reverse((*tab)[i])
	}

	//fmt.Println("EastPole", tab)
	var rockNRoll []string
	for _, col := range *tab {
		rockNRoll = append(rockNRoll, RockNRoll(col))
	}

	//fmt.Println("EastPole", rockNRoll)
	(*tab) = rockNRoll
}

func Day14Star2() {

	lines := importDataByFile()
	/* create tab with columns*/
	var tab []string
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " => ")...)
	}

	cycles := 8

	for i := 0; i < cycles; i++ {
		NorthPole(&tab, lines)
		WestPole(&tab)
		SouthPole(&tab)
		EastPole(&tab)
		if i%100 == 0 {
			fmt.Println(i)
		}
	}

	sum := 0

	for j := 0; j < len(tab); j++ {
		for i := 0; i < len(tab[j]); i++ {
			if string(tab[j][i]) == "O" {
				sum += int(math.Abs(float64(i - len(tab[j]))))
			}
		}
	}

	fmt.Println(sum)

}

func main() {

	startTime := time.Now()
	Day14Star2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
