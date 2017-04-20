package main

import (
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

func main() {
	p9_1()
	p9_2()
	//	p9_3()
}
