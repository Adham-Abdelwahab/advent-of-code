package main

import f "fmt"

func eight() {

	input := `..................................................
................2.................................
......6.........x.0..G............................
..............x5......0..................S........
.....0............................................
..................................y..............e
..........................G...............O.......
.....................0........GO...............d..
.........................8..........e.............
.........6....................................e...
......z6..5...N..x...................eY...........
................6.........5..........Y.E..........
.........X.....N....................E.a...S.....4.
...........................N.2......d.............
...s..................92.....a...................4
............s....................GO........4......
...........................................d.....S
.....................X....N.......................
.........A........................................
.s.....................A....E.......a.........Y...
.....g....s..................E.....Y..............
.............o....................................
...............................3...............O..
.g.................F.3.y..........................
.......F................y.....................d...
..................................X...............
..8....5............X..Z..........................
..g.....8.....na..................................
......................................3...........
.............J.......x............S.Z.............
..2J....h.A...............Z.......................
......A.............................3.............
............J.......n.............................
.8......o....n...........Z........................
..................o..............y................
..F.........................D...............9H....
.................................1.............9..
..................................................
.........h.....n......................f...........
.h....................z..........j.........9......
.......oF............................j............
..........h......z...........7.....1.f............
........................7.......1...H...j........f
........................................f.........
...........................7.......H..............
................................H.................
.............z...........D........................
..............J....................Dj.............
....................................D.............
....................7.......1.....................`

	lines := lines(input)
	max = point{len(lines), len(lines[0])}

	antennas := make(map[rune][]point)
	for x, line := range lines {
		for y, c := range line {
			if c == '.' {
				continue
			}
			if _, ok := antennas[c]; ok {
				antennas[c] = append(antennas[c], point{x, y})
			} else {
				antennas[c] = []point{{x, y}}
			}
		}
	}

	antinodes := make(map[point]bool)
	harmonic_antinodes := make(map[point]bool)

	for _, points := range antennas {
		for index, start := range points {
			for _, end := range points[index+1:] {
				dy := end.y - start.y
				dx := end.x - start.x

				p1 := point{start.x - dx, start.y - dy}
				p2 := point{end.x + dx, end.y + dy}
				if p1.x >= 0 && p1.x < max.x && p1.y >= 0 && p1.y < max.y {
					antinodes[p1] = true
				}
				if p2.x >= 0 && p2.x < max.x && p2.y >= 0 && p2.y < max.y {
					antinodes[p2] = true
				}
			}
		}
	}

	for _, points := range antennas {
		for index, start := range points {
			for _, end := range points[index+1:] {
				dy := end.y - start.y
				dx := end.x - start.x
				p1 := start
				p2 := end

				for {
					p1done := false
					if p1.x >= 0 && p1.x < max.x && p1.y >= 0 && p1.y < max.y {
						harmonic_antinodes[p1] = true
					} else {
						p1done = true
					}
					if p2.x >= 0 && p2.x < max.x && p2.y >= 0 && p2.y < max.y {
						harmonic_antinodes[p2] = true
					} else {
						if p1done {
							break
						}
					}

					p1 = point{p1.x - dx, p1.y - dy}
					p2 = point{p2.x + dx, p2.y + dy}
				}
			}
		}
	}

	f.Printf("8) %v\n", len(antinodes))
	f.Printf("8) %v\n", len(harmonic_antinodes))

}
