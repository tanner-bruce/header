package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: header <text>")
		os.Exit(1)
	}
	str := strings.Join(os.Args[1:], " ")
	left := "## "
	right := " ##"
	l := len(left + right + str)
	header := strings.Repeat("#", l)
	fmt.Println(header)
	fmt.Printf("%s%s%s\n", left, str, right)
	fmt.Println(header)
	os.Exit(0)
}
