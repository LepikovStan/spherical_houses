package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
	"time"
	"strings"
)

func readFile(path string) string {
	var result string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = fmt.Sprintf("%s%s", result, scanner.Text())
	}
	return result
}


var d bool
func initFlags() {
	flag.BoolVar(&d, "d", false, "")
	flag.Parse()
}

func fsm(symbol string, state []int) []int {
	switch symbol {
	case "^":
		newState := []int{state[0], state[1]+1}
		return newState
	case ">":
		newState := []int{state[0]+1, state[1]}
		return newState
	case "<":
		newState := []int{state[0]-1, state[1]}
		return newState
	case "v":
		newState := []int{state[0], state[1]-1}
		return newState
	}

	return state
}

func simple() {
	line := readFile("input.txt")
	state := []int{0,0}
	symbols := strings.Split(line, "")
	housesFields := map[string]int{
		"0,0": 1,
	}
	for _, symbol := range(symbols) {
		state = fsm(symbol, state)
		key := fmt.Sprintf("%d,%d", state[0], state[1])
		if _, ok := housesFields[key]; ok {
			housesFields[key]++
		} else {
			housesFields[key] = 1
		}
	}

	fmt.Println("Result = ", len(housesFields))
}

func update(symbol string, housesFields map[string]int, state []int) (map[string]int, []int) {
	state = fsm(symbol, state)
	key := fmt.Sprintf("%d,%d", state[0], state[1])
	if _, ok := housesFields[key]; ok {
		housesFields[key]++
	} else {
		housesFields[key] = 1
	}
	return housesFields, state
}

func difficult() {
	line := readFile("input.txt")
	stateFirst := []int{0,0}
	stateSecond := []int{0,0}
	symbols := strings.Split(line, "")
	housesFields := map[string]int{
		"0,0": 1,
	}
	for number, symbol := range(symbols) {
		if number%2 == 0 {
			housesFields, stateFirst = update(symbol, housesFields, stateFirst)
		} else {
			housesFields, stateSecond = update(symbol, housesFields, stateSecond)
		}
	}

	fmt.Println("Result = ", len(housesFields))
}

func main() {
	start := time.Now()
	initFlags()
	if d {
		difficult()
	} else {
		simple()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
