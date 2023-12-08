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

type Way struct {
	key   string
	left  string
	right string
}

func Format(ligne string) Way {
	/* split ligne by space and dot*/
	tab := strings.FieldsFunc(ligne, func(r rune) bool { return r == ',' || r == '=' || r == ')' || r == '(' })
	/* remove all blank space*/
	tab = strings.Fields(strings.Join(tab, ""))

	NewWay := Way{key: tab[0], left: tab[1], right: tab[2]}

	return NewWay
}

func FindWay(ligne Way, instruction byte) string {
	if instruction == 'L' {
		return ligne.left
	} else {
		return ligne.right
	}

}

func DayEightStar1() {
	tab := importDataByFile()
	//tab_to_sum := []int{}
	instructions := tab[0]
	var dict_format map[string]Way
	dict_format = make(map[string]Way)
	for _, ligne := range tab[2:] {
		NewWay := Format(ligne)
		dict_format[NewWay.key] = NewWay
	}

	fmt.Println(dict_format)

	i_count := 0
	new_key := "AAA" //Format(tab[2]).key

	final := 0
	for new_key != "ZZZ" {
		if i_count > len(instructions)-1 {
			i_count = 0
		}

		new_key = FindWay(dict_format[new_key], instructions[i_count])
		i_count++
		final++
	}

	fmt.Println(final)

}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func DayEightStar2() {
	tab := importDataByFile()
	//tab_to_sum := []int{}
	instructions := tab[0]
	var dict_format map[string]Way
	dict_format = make(map[string]Way)
	var position []string
	var position2 []string
	for _, ligne := range tab[2:] {
		NewWay := Format(ligne)
		dict_format[NewWay.key] = NewWay

		if NewWay.key[2] == 'A' {
			position = append(position, NewWay.key)
		}
		if NewWay.key[2] == 'Z' {
			position2 = append(position2, NewWay.key)
		}

	}

	fmt.Println(position)

	i_count := 0
	//new_key := "PRA" //Format(tab[2]).key

	results := []int{}

	for _, new_key := range position {
		final := 0

		for new_key[2] != 'Z' {
			if i_count > len(instructions)-1 {
				i_count = 0
			}

			new_key = FindWay(dict_format[new_key], instructions[i_count])
			i_count++
			final++
		}

		fmt.Println(final)
		results = append(results, final)
	}

	fmt.Println(LCM(results[0], results[1], results[2], results[3], results[4], results[5]))

}

func main() {
	startTime := time.Now()
	DayEightStar2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)
}
