package main

import (
	"strings"
	"sync"
)

type point struct {
	x int
	y int
	z int
}

var max point

var wg sync.WaitGroup
var done chan bool

var digits = map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '0': 0}

func lines(input string) []string {
	return strings.Split(input, "\n")
}
