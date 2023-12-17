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

func Direction(l lighting, tab []tile) lighting {
	if l.currentcase.mirror == "/" {
		if l.direction == "up" {
			/* check if the case exists*/
			if l.currentcase.right != -1 {
				l.currentcase = tab[l.currentcase.right]
				tab[l.currentcase.right].visited = true
				l.direction = "right"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "down" {
			if l.currentcase.left != -1 {
				l.currentcase = tab[l.currentcase.left]
				tab[l.currentcase.left].visited = true
				l.direction = "left"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "left" {
			if l.currentcase.down != -1 {
				l.currentcase = tab[l.currentcase.down]
				tab[l.currentcase.down].visited = true
				l.direction = "down"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "right" {
			if l.currentcase.up != -1 {
				l.currentcase = tab[l.currentcase.up]
				tab[l.currentcase.up].visited = true
				l.direction = "up"
			} else {
				l.direction = "out"
			}
		}
	} else if l.currentcase.mirror == "\\" {
		if l.direction == "up" {
			if l.currentcase.left != -1 {
				l.currentcase = tab[l.currentcase.left]
				tab[l.currentcase.left].visited = true
				l.direction = "left"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "down" {
			if l.currentcase.right != -1 {
				l.currentcase = tab[l.currentcase.right]
				tab[l.currentcase.right].visited = true
				l.direction = "right"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "left" {
			if l.currentcase.up != -1 {
				l.currentcase = tab[l.currentcase.up]
				tab[l.currentcase.up].visited = true
				l.direction = "up"
			} else {
				l.direction = "out"
			}
		} else if l.direction == "right" {
			if l.currentcase.down != -1 {
				l.currentcase = tab[l.currentcase.down]
				tab[l.currentcase.down].visited = true
				l.direction = "down"
			} else {
				l.direction = "out"
			}
		}
	} else if l.currentcase.mirror == "-" {
		if l.direction == "left" {
			if l.currentcase.left != -1 {
				l.currentcase = tab[l.currentcase.left]
				tab[l.currentcase.left].visited = true
				l.direction = "left"
			} else {
				l.direction = "out"
			}

		} else if l.direction == "right" {
			if l.currentcase.right != -1 {
				l.currentcase = tab[l.currentcase.right]
				tab[l.currentcase.right].visited = true
				l.direction = "right"
			} else {
				l.direction = "out"
			}
		} else {
			l.direction = "splitdown"
		}
	} else if l.currentcase.mirror == "|" {
		if l.direction == "up" {
			if l.currentcase.up != -1 {
				l.currentcase = tab[l.currentcase.up]
				tab[l.currentcase.up].visited = true

			} else {
				l.direction = "out"
			}
		} else if l.direction == "down" {
			if l.currentcase.down != -1 {
				l.currentcase = tab[l.currentcase.down]
				tab[l.currentcase.down].visited = true

			} else {
				l.direction = "out"
			}
		} else {
			l.direction = "splitup"
		}
	} else {
		if l.direction == "up" && l.currentcase.up != -1 {
			l.currentcase = tab[l.currentcase.up]
			tab[l.currentcase.up].visited = true
		}
		if l.direction == "down" && l.currentcase.down != -1 {
			l.currentcase = tab[l.currentcase.down]
			tab[l.currentcase.down].visited = true
		}
		if l.direction == "left" && l.currentcase.left != -1 {
			l.currentcase = tab[l.currentcase.left]
			tab[l.currentcase.left].visited = true
		}
		if l.direction == "right" && l.currentcase.right != -1 {
			l.currentcase = tab[l.currentcase.right]
			tab[l.currentcase.right].visited = true
		}
	}
	fmt.Println(l)
	return l
}

func LightingPropagation(tab []tile) []tile {
	var lightings []lighting

	/* first lighting*/
	var l lighting
	l.id = 0
	l.currentcase = tab[0]
	l.direction = "right"
	lightings = append(lightings, l)
	z := 0
	for z > 1000 {

		/* faire bouger l'éclair*/
		for i := 0; i < len(lightings); i++ {
			lightings[i] = Direction(lightings[i], tab)

			if lightings[i].direction == "splitup" {
				fmt.Println("splitup")
				var l1 lighting
				l1.id = len(lightings)
				l1.currentcase = lightings[i].currentcase
				l1.direction = "up"
				lightings = append(lightings, l1)
				var l2 lighting
				l2.id = len(lightings)
				l2.currentcase = lightings[i].currentcase
				l2.direction = "down"
				lightings = append(lightings, l2)
				/*remove current lighting*/
				lightings = append(lightings[:i], lightings[i+1:]...)
			} else if lightings[i].direction == "splitleft" {
				var l1 lighting
				l1.id = len(lightings)
				l1.currentcase = lightings[i].currentcase
				l1.direction = "left"
				lightings = append(lightings, l1)
				var l2 lighting
				l2.id = len(lightings)
				l2.currentcase = lightings[i].currentcase
				l2.direction = "right"
				lightings = append(lightings, l2)
				/*remove current lighting*/
				lightings = append(lightings[:i], lightings[i+1:]...)
			} else if lightings[i].direction == "out" {
				lightings = append(lightings[:i], lightings[i+1:]...)
			}
		}
	}
	return tab
}

func Day16Star1() {

	lines := importDataByFile()
	/* create tab with columns*/
	tab := CreateTile(lines)
	tab = LightingPropagation(tab)
	counter := 0
	for _, t := range tab {
		if t.visited {
			counter++
		}
	}
	fmt.Println(counter)
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
