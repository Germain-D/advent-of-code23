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

type tile struct {
	id      int
	mirror  string
	up      int
	down    int
	left    int
	right   int
	visited bool
}

func CreateTile(tab []string) []tile {
	height := len(tab)
	width := len(tab[0])
	var tiles []tile
	counter := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var t tile
			t.id = counter
			t.mirror = string(tab[i][j])
			if i != 0 {
				t.up = counter - width
			}
			if i != height-1 {
				t.down = counter + width
			}
			if j != 0 {
				t.left = counter - 1
			}
			if j != width-1 {
				t.right = counter + 1
			}
			t.visited = false
			tiles = append(tiles, t)
			counter++

		}
	}

	return tiles
}

type lighting struct {
	id          int
	currentcase tile
	direction   string
}

func LightingPropagation(tab []tile) []tile {
	var lightings []lighting
	var tiles []tile

	/* first lighting*/
	var l lighting
	l.id = 0
	l.currentcase = tab[0]
	l.direction = "right"
	lightings = append(lightings, l)

	for len(lightings) > 0 {

		/* faire bouger l'éclair*/

		/* sortie d'un éclair*/
		if lightings[0].direction == "up" && lightings[0].currentcase.up != -1 {
			lightings[0].currentcase = tab[lightings[0].currentcase.up]
		} else if lightings[0].direction == "down" && lightings[0].currentcase.down != -1 {
			lightings[0].currentcase = tab[lightings[0].currentcase.down]
		} else if lightings[0].direction == "left" && lightings[0].currentcase.left != -1 {
			lightings[0].currentcase = tab[lightings[0].currentcase.left]
		} else if lightings[0].direction == "right" && lightings[0].currentcase.right != -1 {
			lightings[0].currentcase = tab[lightings[0].currentcase.right]
		} else {
			lightings = lightings[1:]
		}

	}
	return tiles
}

func Day16Star1() {

	lines := importDataByFile()
	/* create tab with columns*/
	tab := CreateTile(lines)
	fmt.Println(LightingPropagation(tab))

}

func Day16Star2() {

}

func main() {

	startTime := time.Now()
	Day16Star1()
	Day16Star2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
