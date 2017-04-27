package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func p9_1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println()
}

func p9_2() {
	for i, v := range os.Args[1:] {
		fmt.Println(i, v)
	}
	fmt.Println()
}

func p12() {
	// go run c1.go < dummy.txt
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	//p9_1()
	//p9_2()
	p12()
}
