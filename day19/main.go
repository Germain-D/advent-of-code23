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

type Workflow struct {
	id    string
	rules []Rules
	Else  string
}

type Rules struct {
	part   string
	symbol string
	value  int
	wrkfl  string
}

type Part struct {
	x int
	m int
	a int
	s int
}

type workflows map[string]Workflow // map of workflows

func Day19Star1() {

	lines := importDataByFile()
	/* create tab with columns*/
	var workflows workflows
	workflows = make(map[string]Workflow)

	var parts []Part
	switchc := false
	for _, line := range lines {
		if line == "" {
			switchc = true
			continue
		}
		if switchc {
			p := strings.FieldsFunc(line, func(r rune) bool { return r == '{' || r == '}' })
			p = strings.Split(p[0], ",")
			x, err := strconv.Atoi(p[0][2:])
			if err != nil {
				log.Fatal(err)
			}
			m, err := strconv.Atoi(p[1][2:])
			if err != nil {
				log.Fatal(err)
			}
			a, err := strconv.Atoi(p[2][2:])
			if err != nil {
				log.Fatal(err)
			}
			s, err := strconv.Atoi(p[3][2:])
			if err != nil {
				log.Fatal(err)
			}
			parts = append(parts, Part{x: x, m: m, a: a, s: s})
		} else {
			rule := strings.FieldsFunc(strings.Split(line, "{")[1], func(r rune) bool { return r == '{' || r == '}' })
			rule = strings.Split(rule[0], ",")
			var rules []Rules
			for i := 0; i < len(rule)-1; i++ {
				var r Rules
				split := strings.Split(rule[i], ":")
				r.wrkfl = split[1]
				split2 := strings.FieldsFunc(split[0], func(r rune) bool { return r == '<' || r == '>' })
				r.part = split2[0]
				v, err := strconv.Atoi(split2[1])
				if err != nil {
					log.Fatal(err)
				}
				r.value = v
				r.symbol = strings.FieldsFunc(split[0], func(r rune) bool { return r != '<' && r != '>' })[0]

				rules = append(rules, r)
			}

			var w Workflow
			w.id = strings.Split(line, "{")[0]
			w.rules = rules
			w.Else = rule[len(rule)-1]
			workflows[w.id] = w
		}
	}

	sum := 0

	for _, part := range parts {
		start := true
		id := "in"

		for start {
			nomatch := true
			for _, rule := range workflows[id].rules {

				if rule.symbol == "<" {
					if rule.part == "x" {
						if part.x < rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "m" {
						if part.m < rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "a" {
						if part.a < rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "s" {
						if part.s < rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
				} else if rule.symbol == ">" {
					if rule.part == "x" {
						if part.x > rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "m" {
						if part.m > rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "a" {
						if part.a > rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
					if rule.part == "s" {
						if part.s > rule.value {
							id = rule.wrkfl
							nomatch = false
							break
						}
					}
				}

			}
			if nomatch {
				id = workflows[id].Else
			}
			if id == "A" {
				sum += part.x
				sum += part.m
				sum += part.a
				sum += part.s
				start = false
			} else if id == "R" {
				start = false
			}
		}

	}

	fmt.Println(sum)

}

func Calcul(start bool, id string, part Part, workflows workflows, sum *int) {
	for start {
		nomatch := true
		for _, rule := range workflows[id].rules {

			if rule.symbol == "<" {
				if rule.part == "x" {
					if part.x < rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "m" {
					if part.m < rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "a" {
					if part.a < rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "s" {
					if part.s < rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
			} else if rule.symbol == ">" {
				if rule.part == "x" {
					if part.x > rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "m" {
					if part.m > rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "a" {
					if part.a > rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
				if rule.part == "s" {
					if part.s > rule.value {
						id = rule.wrkfl
						nomatch = false
						break
					}
				}
			}

		}
		if nomatch {
			id = workflows[id].Else
		}
		if id == "A" {
			*sum += 1
			start = false
		} else if id == "R" {
			start = false
		}
	}
}

func Day19Star2() {

	lines := importDataByFile()
	/* create tab with columns*/
	var workflows workflows
	workflows = make(map[string]Workflow)

	var parts []Part
	switchc := false
	for _, line := range lines {
		if line == "" {
			switchc = true
			continue
		}
		if switchc {
			p := strings.FieldsFunc(line, func(r rune) bool { return r == '{' || r == '}' })
			p = strings.Split(p[0], ",")
			x, err := strconv.Atoi(p[0][2:])
			if err != nil {
				log.Fatal(err)
			}
			m, err := strconv.Atoi(p[1][2:])
			if err != nil {
				log.Fatal(err)
			}
			a, err := strconv.Atoi(p[2][2:])
			if err != nil {
				log.Fatal(err)
			}
			s, err := strconv.Atoi(p[3][2:])
			if err != nil {
				log.Fatal(err)
			}
			parts = append(parts, Part{x: x, m: m, a: a, s: s})
		} else {
			rule := strings.FieldsFunc(strings.Split(line, "{")[1], func(r rune) bool { return r == '{' || r == '}' })
			rule = strings.Split(rule[0], ",")
			var rules []Rules
			for i := 0; i < len(rule)-1; i++ {
				var r Rules
				split := strings.Split(rule[i], ":")
				r.wrkfl = split[1]
				split2 := strings.FieldsFunc(split[0], func(r rune) bool { return r == '<' || r == '>' })
				r.part = split2[0]
				v, err := strconv.Atoi(split2[1])
				if err != nil {
					log.Fatal(err)
				}
				r.value = v
				r.symbol = strings.FieldsFunc(split[0], func(r rune) bool { return r != '<' && r != '>' })[0]

				rules = append(rules, r)
			}

			var w Workflow
			w.id = strings.Split(line, "{")[0]
			w.rules = rules
			w.Else = rule[len(rule)-1]
			workflows[w.id] = w
		}
	}

	sum := 0
	for x := 100; x > 0; x-- {
		for m := 100; m > 0; m-- {
			for a := 100; a > 0; a-- {
				for s := 100; s > 0; s-- {
					var part Part
					part.x = x
					part.m = m
					part.a = a
					part.s = s
					start := true
					id := "in"
					Calcul(start, id, part, workflows, &sum)

				}
			}
		}
	}

	fmt.Println(sum)

}

func main() {

	startTime := time.Now()
	Day19Star1()
	Day19Star2()
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
