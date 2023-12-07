package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*doesn't work (have to be connected)*/
func importDataByUrl() []string {
	fmt.Println("Importing data...")
	resp, err := http.Get("https://adventofcode.com/2023/day/1/input")
	if err != nil {
		panic(err)
	}

	/* create tab with rows*/
	var tab []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}

	defer resp.Body.Close()
	fmt.Println("Done.")
	return tab
}

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

func GetDigits(ligne string) []string {
	var digits []string
	for _, i := range ligne {
		if unicode.IsNumber(i) {
			digits = append(digits, string(i))
		}

	}
	return digits
}

func DayOneStar1() {
	tab := importDataByFile()
	tab_to_sum := []int{}
	for _, ligne := range tab {
		res := GetDigits(ligne)
		final := res[0] + res[len(res)-1]
		i, err := strconv.Atoi(final)
		if err != nil {
			panic(err)
		}
		tab_to_sum = append(tab_to_sum, i)
	}
	var sum int
	for _, i := range tab_to_sum {
		sum += i
	}
	fmt.Println(sum)

}

func GetDigits2(ligne string) int {

	type Number struct {
		valeur   string
		position int
	}

	var min Number
	var max Number

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	count := 1

	for _, i := range numbers {
		if strings.Contains(ligne, i) {

			for j := 0; j < len(ligne); j++ {
				index := strings.Index(string(ligne)[j:], i)
				if index == -1 {
					break
				}
				if min == (Number{}) {
					min = Number{strconv.Itoa(count), index}
				} else {
					if index < min.position {
						min = Number{strconv.Itoa(count), index}
					}
				}
				if max == (Number{}) {
					max = Number{strconv.Itoa(count), index}
				} else {
					if index > max.position {
						max = Number{strconv.Itoa(count), index}
					}
				}
				j += index
			}

		}
		count++
	}

	position := 0
	for _, i := range ligne {
		if unicode.IsNumber(i) {
			if min == (Number{}) {
				min = Number{string(i), position}
			} else {
				if position < min.position {
					min = Number{string(i), position}
				}
			}
			if max == (Number{}) {
				max = Number{string(i), position}
			} else {
				if position > max.position {
					max = Number{string(i), position}
				}
			}
		}
		position++
	}

	final := min.valeur + max.valeur
	i, err := strconv.Atoi(final)
	if err != nil {
		panic(err)
	}

	return i

}

func GetDigits3(ligne string) int {

	type Number struct {
		valeur   string
		position int
	}

	var min Number
	var max Number

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	count := 1

	for _, i := range numbers {
		ligne = strings.Replace(ligne, i, i+strconv.Itoa(count)+i, -1)
		count++
	}

	position := 0
	for _, i := range ligne {
		if unicode.IsNumber(i) {
			if min == (Number{}) {
				min = Number{string(i), position}
			} else {
				if position < min.position {
					min = Number{string(i), position}
				}
			}
			if max == (Number{}) {
				max = Number{string(i), position}
			} else {
				if position > max.position {
					max = Number{string(i), position}
				}
			}
		}
		position++
	}

	final := min.valeur + max.valeur
	fmt.Println(final)
	i, err := strconv.Atoi(final)
	if err != nil {
		panic(err)
	}

	return i

}

func DayOneStar2() {
	tab := importDataByFile()
	tab_to_sum := []int{}
	for _, ligne := range tab {
		i := GetDigits3(ligne)
		tab_to_sum = append(tab_to_sum, i)
	}
	fmt.Println(tab_to_sum)
	var sum int
	for _, i := range tab_to_sum {
		sum += i
	}
	fmt.Println(sum)

}

func main() {
	DayOneStar2()
}
