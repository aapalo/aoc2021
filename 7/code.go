package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

var (
	// 1, 2 or 3
	//part int = 2
	//sample bool = true
)

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func partone(s []string) (ans int) {
	ans = 1
	crabs_s := strings.Split(s[0], ",")
	var crabs = []int{}
	for _, i := range(crabs_s) {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, j)
	}
	//fmt.Println(crabs)
	var fuels = []int{}
	_, maxcrab := MinMax(crabs)
	fuel := 0
	for i := 0; i < maxcrab; i++ {
		fuel = 0
		for _, c := range(crabs) {
			dist := c - i
			if dist < 0 {
				dist = -dist
			}
			fuel += dist
		}
		//fmt.Println(i, fuel)
		fuels = append(fuels, fuel)
	}
	minfuel, _ := MinMax(fuels)
	//fmt.Println("Part one:", ans)
	return minfuel
}

func countfuel(steps int) (f int) {
	f = 0
	for i := 1; i <= steps; i++ {
		f += i
	}
	return f
}

func parttwo(s []string) (ans int) {
	ans = 1
	crabs_s := strings.Split(s[0], ",")
	var crabs = []int{}
	for _, i := range(crabs_s) {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, j)
	}
	//fmt.Println(crabs)
	var fuels = []int{}
	_, maxcrab := MinMax(crabs)
	fuel := 0
	for i := 0; i < maxcrab; i++ {
		fuel = 0
		for _, c := range(crabs) {
			dist := c - i
			if dist < 0 {
				dist = -dist
			}
			fuel += countfuel(dist)
		}
		//fmt.Println(i, fuel)
		fuels = append(fuels, fuel)
	}
	minfuel, _ := MinMax(fuels)
	return minfuel
}

func main() {
	var filename string = "./sample.txt" //"./input.txt"
	notsample := 0
	part := 3
	if len(os.Args) > 1 {
		part, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Println("Run with 'go run code.go <parts> <notsample>'")
	}
	if len(os.Args) > 2 {
		notsample, _ = strconv.Atoi(os.Args[2])
	}
	if notsample != 0 {
		filename = "./input.txt"
	}
	fmt.Println("Part:", part, filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var s []string
	var t string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t = scanner.Text()
		s = append(s, t)
		//fmt.Println(s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//var fuelsum int = 0
	if part == 1 {
		fmt.Println("Part one:", partone(s))
	} else if part == 2 {
		fmt.Println("Part two:", parttwo(s))
	} else {
		fmt.Println("Part one:", partone(s))
		fmt.Println("Part two:", parttwo(s))
	}
}
