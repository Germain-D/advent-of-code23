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

func tabstringcontains(elems []string, v string) bool {
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

/* flipflop % -> base off, ignore high pulse, flip on low pulse */
/* conjonction & -> base low pulse, remember pulse received by each input, send low pulse if all inputs are high, else send high */
/* broadcast -> send pulse to all outputs */
/* button -> send pulse when pressed to broadcaster*/

/* 0 -> off,low , 1 -> on,high */

type Module struct {
	name    string
	state   bool
	symbol  string
	inputs  []Pulse
	outputs []string
}

type Pulse struct {
	state    string
	sender   string
	receiver string
}

func flipflop(pulse string, flipstate bool) bool {
	if pulse == "low" {
		return !flipstate
	}
	return flipstate
}

func conjonction(pulse Pulse, conj Module) (bool, Module) {
	// fmt.Println("conjonction")
	// // fmt.Println(pulse)
	// fmt.Println(conj.inputs)

	for i, input := range conj.inputs {
		if input.sender == pulse.sender {
			/* change state of input */
			conj.inputs[i].state = pulse.state

		}
	}
	for _, input := range conj.inputs {
		if input.state == "low" {
			return true, conj
		}
	}
	return false, conj
}

type modules map[string]Module

func Day20Star1(lines []string) {
	var modules modules
	modules = make(map[string]Module)
	for _, line := range lines {
		l := strings.Split(line, " -> ")
		module := l[0]
		outputs := strings.Split(l[1], ", ")
		if strings.Contains(module, "%") {
			modules[strings.Split(module, "%")[1]] = Module{strings.Split(module, "%")[1], false, "%", nil, outputs}
		} else if strings.Contains(module, "&") {
			modules[strings.Split(module, "&")[1]] = Module{strings.Split(module, "&")[1], false, "&", nil, outputs}
		} else {
			modules[module] = Module{module, false, "broad", nil, outputs}
		}
	}

	/* on parcours tous les modules pour set les inputs des conjonctions */
	for _, mod := range modules {
		if mod.symbol == "&" {
			for _, mod2 := range modules {
				if tabstringcontains(mod2.outputs, mod.name) {
					input := modules[mod.name].inputs
					input2 := Pulse{"low", mod2.name, mod.name}
					input = append(input, input2)
					modules[mod.name] = Module{mod.name, false, "&", input, mod.outputs}

				}
			}

		}
	}

	fmt.Println(modules)

	var pulses_sent []Pulse
	button_pressed := 1000
	pulse_counter := 0
	for i := 0; i < button_pressed; i++ {
		var attente []string
		pulse := Pulse{"low", "button", "broadcaster"}
		pulses_sent = append(pulses_sent, pulse) /* button pressed send low*/
		pulse_counter += 1
		fmt.Printf("%s -%s-> %s\n", "button", pulse.state, "broadcaster")
		for _, out := range modules["broadcaster"].outputs {
			attente = append(attente, out)
			pulse = Pulse{"low", "broadcaster", out}
			pulses_sent = append(pulses_sent, pulse) /* broadcaster send low*/
			fmt.Printf("%s -%s-> %s\n", "broadcaster", pulse.state, out)

		}

		for len(attente) > 0 {
			base_attente := attente
			attente = nil
			for _, att := range base_attente {
				if att != "output" {
					var mod Module
					mod = modules[att]
					pulse := pulses_sent[pulse_counter]

					if mod.symbol == "%" {
						/* do flipflop */
						if pulse.state != "high" {
							new_state := flipflop(pulse.state, mod.state)
							modules[att] = Module{mod.name, new_state, mod.symbol, mod.inputs, mod.outputs}
							if new_state == true {
								for _, out := range mod.outputs {
									attente = append(attente, out)
									pulse = Pulse{"high", att, out}
									pulses_sent = append(pulses_sent, pulse)
									fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
								}
							} else {
								for _, out := range mod.outputs {
									attente = append(attente, out)
									pulse = Pulse{"low", att, out}
									pulses_sent = append(pulses_sent, pulse)
									fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
								}
							}
						}
					} else if mod.symbol == "&" {
						/* do conjonction */
						new_state, mod := conjonction(pulse, mod)
						modules[att] = mod
						if new_state == true {
							for _, out := range mod.outputs {
								attente = append(attente, out)
								pulse = Pulse{"high", att, out}
								pulses_sent = append(pulses_sent, pulse)
								fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
							}
						} else {
							for _, out := range mod.outputs {
								attente = append(attente, out)
								pulse = Pulse{"low", att, out}
								pulses_sent = append(pulses_sent, pulse)
								fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
							}
						}
					} else {
						/* do broadcast */

					}
				}
				pulse_counter += 1
			}
		}
	}

	low := 0
	high := 0
	for _, pulse := range pulses_sent {
		if pulse.state == "low" {
			low += 1
		} else {
			high += 1
		}
	}
	/*17 low, 11 high*/
	fmt.Printf("low: %d, high: %d\n", low, high)
	fmt.Println(low * high)

}

func Day20Star2(lines []string) {

	var modules modules
	modules = make(map[string]Module)
	for _, line := range lines {
		l := strings.Split(line, " -> ")
		module := l[0]
		outputs := strings.Split(l[1], ", ")
		if strings.Contains(module, "%") {
			modules[strings.Split(module, "%")[1]] = Module{strings.Split(module, "%")[1], false, "%", nil, outputs}
		} else if strings.Contains(module, "&") {
			modules[strings.Split(module, "&")[1]] = Module{strings.Split(module, "&")[1], false, "&", nil, outputs}
		} else {
			modules[module] = Module{module, false, "broad", nil, outputs}
		}
	}

	/* on parcours tous les modules pour set les inputs des conjonctions */
	for _, mod := range modules {
		if mod.symbol == "&" {
			for _, mod2 := range modules {
				if tabstringcontains(mod2.outputs, mod.name) {
					input := modules[mod.name].inputs
					input2 := Pulse{"low", mod2.name, mod.name}
					input = append(input, input2)
					modules[mod.name] = Module{mod.name, false, "&", input, mod.outputs}

				}
			}

		}
	}

	fmt.Println(modules)

	button_pressed := 100000000

	for i := 0; i < button_pressed; i++ {
		var pulses_sent []Pulse
		pulse_counter := 0
		var attente []string
		pulse := Pulse{"low", "button", "broadcaster"}
		pulses_sent = append(pulses_sent, pulse) /* button pressed send low*/
		pulse_counter += 1
		//fmt.Printf("%s -%s-> %s\n", "button", pulse.state, "broadcaster")
		for _, out := range modules["broadcaster"].outputs {
			attente = append(attente, out)
			pulse = Pulse{"low", "broadcaster", out}
			pulses_sent = append(pulses_sent, pulse) /* broadcaster send low*/
			//fmt.Printf("%s -%s-> %s\n", "broadcaster", pulse.state, out)

		}

		for len(attente) > 0 {
			base_attente := attente
			attente = nil
			for _, att := range base_attente {
				if att != "output" {
					var mod Module
					mod = modules[att]
					pulse := pulses_sent[pulse_counter]
					if pulse.receiver == "rx" && pulse.state == "low" {
						fmt.Println("ICI")
						fmt.Println(i)
						i = button_pressed
						return
					}
					if mod.symbol == "%" {
						/* do flipflop */
						if pulse.state != "high" {
							new_state := flipflop(pulse.state, mod.state)
							modules[att] = Module{mod.name, new_state, mod.symbol, mod.inputs, mod.outputs}
							if new_state == true {
								for _, out := range mod.outputs {
									attente = append(attente, out)
									pulse = Pulse{"high", att, out}
									pulses_sent = append(pulses_sent, pulse)
									//fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
								}
							} else {
								for _, out := range mod.outputs {
									attente = append(attente, out)
									pulse = Pulse{"low", att, out}
									pulses_sent = append(pulses_sent, pulse)
									//fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
								}
							}
						}
					} else if mod.symbol == "&" {
						/* do conjonction */
						new_state, mod := conjonction(pulse, mod)
						modules[att] = mod
						if new_state == true {
							for _, out := range mod.outputs {
								attente = append(attente, out)
								pulse = Pulse{"high", att, out}
								pulses_sent = append(pulses_sent, pulse)
								//fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
							}
						} else {
							for _, out := range mod.outputs {
								attente = append(attente, out)
								pulse = Pulse{"low", att, out}
								pulses_sent = append(pulses_sent, pulse)
								//fmt.Printf("%s -%s-> %s\n", att, pulse.state, out)
							}
						}
					} else {
						/* do broadcast */

					}
				}
				pulse_counter += 1
			}
		}
	}

}

func main() {

	startTime := time.Now()
	lines := importDataByFile()
	Day20Star1(lines)
	Day20Star2(lines)
	endTime := time.Since(startTime)
	fmt.Printf("Time taken in Go: %s\n", endTime)

}
