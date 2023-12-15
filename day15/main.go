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

	fmt.Println(score)
	return score
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Hash(s string) int64 {
	hash := int64(0)
	for _, c := range s {
		hash += int64(c)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

func Day15Star1() {

	lines := importDataByFile()
	/* create tab with columns*/
	var tab []string
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ",")...)
	}

	sum := 0
	for _, elem := range tab {
		sum += int(Hash(elem))
	}

	fmt.Println(sum)

}

type Lens struct {
	id    string
	power int
}

type Box struct {
	id  int
	tab []Lens
}

func Hash2(s string) int64 {
	var label string
	if strings.Contains(s, "-") {
		label = strings.Split(s, "-")[0]
	} else {
		label = strings.Split(s, "=")[0]
	}

	hash := int64(0)
	for _, c := range label {
		hash += int64(c)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

func Day15Star2() {

	lines := importDataByFile()

	/* create 256 boxes*/
	var boxes []Box
	for i := 0; i < 256; i++ {
		var box Box
		box.id = i
		boxes = append(boxes, box)
	}

	/* create tab with columns*/
	var tab []string
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ",")...)
	}

	for _, elem := range tab {
		box_id := int(Hash2(elem))
		var label string
		if strings.Contains(elem, "-") {
			label = strings.Split(elem, "-")[0]
			for i := 0; i < len(boxes[box_id].tab); i++ {
				if boxes[box_id].tab[i].id == label {
					boxes[box_id].tab = append(boxes[box_id].tab[:i], boxes[box_id].tab[i+1:]...)
				}
			}

		} else {
			label = strings.Split(elem, "=")[0]
			power, err := strconv.Atoi(strings.Split(elem, "=")[1])
			if err != nil {
				log.Fatal(err)
			}
			powchange := false
			for i := 0; i < len(boxes[box_id].tab); i++ {
				if boxes[box_id].tab[i].id == label {
					boxes[box_id].tab[i].power = power
					powchange = true
				}
			}
			if powchange == false {
				boxes[box_id].tab = append(boxes[box_id].tab, Lens{id: label, power: power})
			}

		}

	}
	res := 0
	for _, box := range boxes {
		if len(box.tab) > 0 {
			for i := 0; i < len(box.tab); i++ {
				res += (box.id + 1) * (i + 1) * box.tab[i].power
			}
		}
	}
	fmt.Println(res)

}

func main() {

	startTime := time.Now()
	Day15Star1()
	Day15Star2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
