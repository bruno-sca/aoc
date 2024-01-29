package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed test.txt
var input string

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", SolveA(input))
	fmt.Println("Execution Time (Part 1):", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", SolveB(input))
	fmt.Println("Execution Time (Part 2):", time.Since(start))
}
