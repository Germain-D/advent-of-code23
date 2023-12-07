package main

import (
	"fmt"
	"log"
	"os"
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

func isPossible(ligne string) int {
	type Cube struct {
		number int
		color  string
	}

	bag := []Cube{{number: 12, color: "red"}, {number: 14, color: "blue"}, {number: 13, color: "green"}}
	my_bag := []Cube{}

	/* split ligne by space and dot*/
	tab := strings.FieldsFunc(ligne, func(r rune) bool { return r == ',' || r == ';' || r == ':' })
	id := strings.Split(tab[0], " ")[1]

	for _, cube := range tab[1:] {
		/* split cube by space*/
		cube_split := strings.Split(cube, " ")
		/* get number and color*/
		number, err := strconv.Atoi(cube_split[1])
		if err != nil {
			panic(err)
		}
		color := cube_split[2]

		/* if color is in my_bag, add number to the number of the cube*/
		/*
			iscolor := false
			for i := 0; i < len(my_bag); i++ {
				if my_bag[i].color == color {
					my_bag[i].number += number
					iscolor = true
				}
			}*/
		/* if color is not in my_bag, create a new cube*/
		/*
			if iscolor == false {
				cube := Cube{number: number, color: color}
				my_bag = append(my_bag, cube)
			}*/
		cube := Cube{number: number, color: color}
		my_bag = append(my_bag, cube)

	}

	for _, cube := range my_bag {
		for _, cube2 := range bag {
			if cube.color == cube2.color {
				if cube.number > cube2.number {
					return 0
				}
			}
		}
	}

	fmt.Println(my_bag)

	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	return i
}

func DayTwoStar1() {
	tab := importDataByFile()
	tab_to_sum := []int{}
	for _, ligne := range tab {
		res := isPossible(ligne)
		tab_to_sum = append(tab_to_sum, res)
	}
	var sum int
	for _, i := range tab_to_sum {
		sum += i
	}
	fmt.Println(sum)

}

func isPossible2(ligne string) int {
	type Cube struct {
		number int
		color  string
	}

	my_bag := []Cube{}

	/* split ligne by space and dot*/
	tab := strings.FieldsFunc(ligne, func(r rune) bool { return r == ',' || r == ';' || r == ':' })

	for _, cube := range tab[1:] {
		/* split cube by space*/
		cube_split := strings.Split(cube, " ")
		/* get number and color*/
		number, err := strconv.Atoi(cube_split[1])
		if err != nil {
			panic(err)
		}
		color := cube_split[2]

		/* if color is in my_bag, add number to the number of the cube*/
		iscolor := false
		for i := 0; i < len(my_bag); i++ {
			if my_bag[i].color == color {
				iscolor = true
				if my_bag[i].number < number {
					my_bag[i].number = number

				}
			}
		}
		/* if color is not in my_bag, create a new cube*/

		if iscolor == false {
			cube := Cube{number: number, color: color}
			my_bag = append(my_bag, cube)
		}

	}

	power := 1

	for _, cube := range my_bag {
		power = cube.number * power
	}
	fmt.Println(power)

	fmt.Println(my_bag)

	return power
}

func DayTwoStar2() {
	tab := importDataByFile()
	tab_to_sum := []int{}
	for _, ligne := range tab {
		res := isPossible2(ligne)
		tab_to_sum = append(tab_to_sum, res)
	}
	var sum int
	for _, i := range tab_to_sum {
		sum += i
	}
	fmt.Println(sum)

}

func main() {
	DayTwoStar2()
}
