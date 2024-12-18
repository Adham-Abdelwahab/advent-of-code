package main

import (
	"strings"
	"sync"
)

type point struct {
	x int
	y int
}

var max point

var wg sync.WaitGroup
var done chan bool

func lines(input string) []string {
	return strings.Split(input, "\n")
}
