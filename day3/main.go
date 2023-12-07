package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

type PosNumber struct {
	StartPos  int
	EndPos    int
	LineIndex int
	Number    int
}

func findNumbers(str string) []*PosNumber {
	// Define the regular expression for finding numbers
	re := regexp.MustCompile(`\d+`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatchIndex(str, -1)

	// Extract numbers along with start and end positions
	result := make([]*PosNumber, len(matches))
	for i, match := range matches {
		start := match[0]
		end := match[1]
		number, _ := strconv.Atoi(str[start:end])

		result[i] = &PosNumber{
			Number:   number,
			StartPos: start,
			EndPos:   end,
		}
	}

	return result
}

func symbolAround(number *PosNumber, lines []string) bool {
	from := number.StartPos - 1
	if from < 0 {
		from = 0
	}
	to := number.EndPos + 1
	if to > len(lines[0]) {
		to = len(lines[0])
	} // assume all lines have same len

	// loop three lines
	for looplines := number.LineIndex - 1; looplines <= number.LineIndex+1; looplines++ {
		if looplines < 0 || looplines >= len(lines) {
			continue
		}
		// inspect line characters
		symbolFound := strings.IndexAny(lines[looplines][from:to], "+#$*@/=%-&")
		// we know enough already
		if symbolFound > -1 {
			return true
		}
	}

	return false
}

func EngineCheck1(top string, mid string, bottom string) int {
	/* split if not alphanumeric*/
	//pattern := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	//top_split := pattern.Split(top, -1)
	//mid_split := pattern.Split(mid, -1)
	//bottom_split := pattern.Split(bottom, -1)
	//fmt.Println(mid)
	//fmt.Println(mid_split)
	for i := 0; i < len(mid); i++ {
		/* if char is alphanumeric, check if it's a tree*/
		if mid[i] != '.' {
			/* if it's a tree, check if the char above is a tree*/
			if top[i] != '.' {
				/* if it's a tree, check if the char above is a tree*/
				if bottom[i] != '.' {
					/* if it's a tree, return 1*/
					return 1
				}
			}
		}
	}

	return 0
}

func DayThreeStar1() {
	// tab := importDataByFile()
	// tab_to_sum := []int{}
	// /*for _, ligne := range tab {
	// 	res := EngineCheck1(ligne)
	// 	tab_to_sum = append(tab_to_sum, res)
	// }*/

	// match, _ := regexp.MatchString("[^a-zA-Z0-9.]", "999999")
	// fmt.Println(match)
	// pattern := regexp.MustCompile(`[^a-zA-Z0-9.]`)
	// for i := 0; i < len(tab); i++ {
	// 	if pattern.MatchString(tab[i]) {
	// 		top := tab[i-1]
	// 		mid := tab[i]
	// 		bottom := tab[i+1]
	// 		EngineCheck1(top, mid, bottom)
	// 	}
	// }

	// var sum int
	// for _, i := range tab_to_sum {
	// 	sum += i
	// }
	// fmt.Println(sum)
	lines := importDataByFile()

	var allNumbers []*PosNumber
	for lineIndex, line := range lines {
		numbersForLine := findNumbers(line)

		fmt.Printf("%s\n", line)

		for _, res := range numbersForLine {
			// record the lineIndex for found numbers
			res.LineIndex = lineIndex
			fmt.Printf("line: %d, cols:(%d-%d), n:%d\n ", res.LineIndex, res.StartPos, res.EndPos, res.Number)
		}
		allNumbers = append(allNumbers, numbersForLine...)
	}

	// Now we have all numbers and their locations in allNumbers.
	// When we inspect the surrounding for symbols we should
	// be able to find out what numbers are valid
	// to use those in the totalSum

	totalSum := 0
	for _, number := range allNumbers {
		if symbolAround(number, lines) {
			totalSum += number.Number
		}
	}

	fmt.Printf("totalSum: %d\n", totalSum)

}

func main() {
	DayThreeStar1()
}
