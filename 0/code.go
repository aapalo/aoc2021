package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	//"strings"
)

var (
	// 1, 2 or 3
	//part int = 3
	//sample bool = false
)

func partone(s []string) (ans int) {
	ans = 1
	//fmt.Println("Part one:", ans)
	return ans
}

func parttwo(s []string) (ans int) {
	ans = 2
	//mt.Println("Part two:", ans)
	return ans
}

func main() {
	var filename string = "./sample.txt" //"./input.txt"
	notsample := 0
	part := 3
	if len(os.Args) > 1 {
		part, _ = strconv.Atoi(os.Args[1])
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
