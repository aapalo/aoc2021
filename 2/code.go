package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	//"strconv"
)

var (
	// 1, 2 or 3
	part int = 1
	sample bool = true
)

func helper(val int) int {
	if sample {
		return 1
	} else {
		return 0
	}
}

type Submarine struct {
	depth int
	position int
	aim int
}

func partone(s []string) (ans int) {
	ans = 1
	depth := 0
	position := 0
	//fmt.Println(s)
	for _, row := range s {
		//strconv.Atoi(row)
		fmt.Println(row[1])//, task, num)
	}
	ans = depth * position
	//fmt.Println("Part one:", ans)
	return ans
}

func parttwo(s []string) (ans int) {
	ans = 2
	//mt.Println("Part two:", ans)
	return ans
}

func main() {
	var filename string = "./input.txt"
	if sample {
		filename = "./sample.txt"
	}
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
		fmt.Println(s)
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
