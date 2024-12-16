package main

import "strings"

type point struct {
	x int
	y int
}

func lines(input string) []string {
	return strings.Split(input, "\n")
}
