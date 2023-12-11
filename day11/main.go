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

func ExpandRows(universe []string) []string {
	var rowtab []int
	rowadd := strings.Repeat(".", len(universe[0]))

	for i := 0; i < len(universe); i++ {
		if !strings.Contains(universe[i], "#") {
			/* add row */
			rowtab = append(rowtab, i)

		}
	}

	for j := 0; j < len(rowtab); j++ {
		universe = append(universe[:rowtab[j]], append([]string{rowadd}, universe[rowtab[j]:]...)...)
		/* add 1 to all rowtab */
		for k := j + 1; k < len(rowtab); k++ {
			rowtab[k]++
		}
	}

	return universe
}

func ExpandColumns(universe []string) []string {
	var columntab []int
	columnadd := "."

	for i := 0; i < len(universe[0]); i++ {
		if !strings.Contains(universe[0][i:], "#") {
			/* add column */
			columntab = append(columntab, i)

		}
	}

	/* add one column at the column tab index */
	for _, index := range columntab {
		for j := 0; j < len(universe); j++ {
			universe[j] = universe[j][:index] + columnadd + universe[j][index:]
		}
	}

	return universe
}

func ExpandUniverse(universe []string) []string {
	universe = ExpandRows(universe)
	universe = ExpandColumns(universe)
	return universe
}

func PrintUniverse(universe []string) {
	for i := 0; i < len(universe); i++ {
		fmt.Println(universe[i])
	}
}

type Case struct {
	id          int
	neighbours  []int
	neighbours2 map[int]int
	isGalaxy    bool
}

func GetNeighbours(universe []string) []Case {
	fmt.Println(len(universe), len(universe[0]))
	PrintUniverse(universe)
	var universe2 []Case
	idcount := 0
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[0]); j++ {
			if universe[i][j] == '#' {
				universe2 = append(universe2, Case{id: idcount, neighbours: []int{}, isGalaxy: true})
				idcount++
			} else {
				universe2 = append(universe2, Case{id: idcount, neighbours: []int{}, isGalaxy: false})
				idcount++
			}
		}
	}

	idcount = 0
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			/*if corner */
			/* max 4 neighbours */
			if i == 0 && j == 0 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
			} else if i == 0 && j == len(universe[i])-1 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
			} else if i == len(universe)-1 && j == 0 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			} else if i == len(universe)-1 && j == len(universe[i])-1 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			} else if i == 0 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
			} else if i == len(universe)-1 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			} else if j == 0 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			} else if j == len(universe[i])-1 {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			} else {
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+1)
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount+len(universe[0]))
				universe2[idcount].neighbours = append(universe2[idcount].neighbours, idcount-len(universe[0]))
			}
			idcount++
		}
	}

	fmt.Println(universe2)
	return universe2
}

type GalaxyPairs struct {
	a int
	b int
}

func GetGalaxyPairs(universe []Case) []GalaxyPairs {
	var galaxyP []GalaxyPairs
	var galaxy []Case

	for i := 0; i < len(universe); i++ {
		if universe[i].isGalaxy {
			galaxy = append(galaxy, universe[i])
		}
	}
	fmt.Println("--------------------")
	fmt.Println(galaxy)
	fmt.Println("--------------------")

	exist := false
	for i := 0; i < len(galaxy); i++ {
		for j := i + 1; j < len(galaxy); j++ {
			for _, pairs := range galaxyP {
				if pairs.a == galaxy[i].id && pairs.b == galaxy[j].id {
					exist = true
				} else if pairs.a == galaxy[j].id && pairs.b == galaxy[i].id {
					exist = true
				}
			}
			if !exist {
				galaxyP = append(galaxyP, GalaxyPairs{a: galaxy[i].id, b: galaxy[j].id})
			}

		}
	}

	fmt.Println(len(galaxyP))

	return galaxyP
}
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
func ShortestPath(universe []Case, a int, b int) int {
	start := universe[a]
	end := universe[b]
	for i := 0; i < len(start.neighbours); i++ {
		if Contains(start.neighbours, end.id) {
			return 1
		}
	}

}*/

func dijkstra(graph map[int]Case, startID, targetID int) []int {
	// Initialiser les distances avec l'infini sauf pour le nœud de départ, qui a une distance de 0
	distances := make(map[int]int)
	for id := range graph {
		distances[id] = math.MaxInt32
	}
	distances[startID] = 0

	// Initialiser un ensemble pour suivre les nœuds visités
	visited := make(map[int]bool)

	// Boucle principale
	for len(visited) < len(graph) {
		// Trouver le nœud non visité avec la distance minimale actuelle
		currentID := findMinDistanceNode(distances, visited)

		// Marquer le nœud actuel comme visité
		visited[currentID] = true

		// Mettre à jour les distances des voisins du nœud actuel
		for neighborID, edgeWeight := range graph[currentID].neighbours2 {
			if !visited[neighborID] {
				newDistance := distances[currentID] + edgeWeight
				if newDistance < distances[neighborID] {
					distances[neighborID] = newDistance
				}
			}
		}
	}

	// Reconstruction du chemin le plus court
	shortestPath := []int{targetID}
	currentID := targetID

	for currentID != startID {
		for neighborID, edgeWeight := range graph[currentID].neighbours2 {
			if distances[currentID]-edgeWeight == distances[neighborID] {
				shortestPath = append([]int{neighborID}, shortestPath...)
				currentID = neighborID
				break
			}
		}
	}

	return shortestPath
}

func findMinDistanceNode(distances map[int]int, visited map[int]bool) int {
	minDistance := math.MaxInt32
	var minID int

	for id, distance := range distances {
		if !visited[id] && distance < minDistance {
			minDistance = distance
			minID = id
		}
	}

	return minID
}
func DayElevenStar1() {
	tab := importDataByFile()
	//tab_to_sum := []int{}

	/* print row and column size */
	fmt.Println("Before")
	fmt.Println(len(tab), len(tab[0]))

	tab = ExpandUniverse(tab)

	/* print row and column size */
	fmt.Println("After")
	fmt.Println(len(tab), len(tab[0]))

	universe := GetNeighbours(tab)
	galaxy := GetGalaxyPairs(universe)

	fmt.Println(galaxy)

	/*sum := 0
	for _, pairs := range galaxy {
		sum += ShortestPath(universe, pairs.a, pairs.b)
	}

	fmt.Println(sum)*/
	graph := make(map[int]Case)
	for _, node := range universe {
		node.neighbours2 = make(map[int]int)
		for _, neighbourID := range node.neighbours {
			node.neighbours2[neighbourID] = 1
		}
		graph[node.id] = node
	}

	sum := 0
	for _, pairs := range galaxy {
		startID := pairs.a
		endID := pairs.b

		shortestPath := dijkstra(graph, startID, endID)

		sum += len(shortestPath)

	}

	fmt.Println(sum)

}

func main() {
	startTime := time.Now()
	DayElevenStar1()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)
}
