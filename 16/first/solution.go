package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(inputs, "\n")

	var mapp [][]rune
	var mapEnergy [][]bool

	for _, row := range rows {
		var newRow []rune
		var newRowBool []bool
		for _, char := range []rune(row) {
			newRow = append(newRow, char)
			newRowBool = append(newRowBool, false)
		}
		mapp = append(mapp, newRow)
		mapEnergy = append(mapEnergy, newRowBool)
	}

	startCoord := Coords{0, 0}
	startDir := Right

	calcBeam(&mapp, &mapEnergy, startCoord, startDir)

	res := 0
	for _, row := range mapEnergy {
		for _, cell := range row {
			if cell {
				res++
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func calcBeam(mapp *[][]rune, energy *[][]bool, coord Coords, dir Direction) {
	var coords []Coords
	coords = append(coords, coord)
	var directions []Direction
	directions = append(directions, dir)

	var mapUp [][]bool
	var mapDown [][]bool
	var mapRight [][]bool
	var mapLeft [][]bool

	for _, row := range *mapp {
		var newUp []bool
		var newDown []bool
		var newRight []bool
		var newLeft []bool

		for _ = range row {
			newUp = append(newUp, false)
			newDown = append(newDown, false)
			newRight = append(newRight, false)
			newLeft = append(newLeft, false)
		}
		mapUp = append(mapUp, newUp)
		mapDown = append(mapDown, newDown)
		mapRight = append(mapRight, newRight)
		mapLeft = append(mapLeft, newLeft)
	}

	ind := 0

	for ind < len(coords) {
		coord = coords[ind]
		dir = directions[ind]
		for {
			if coord.row < 0 || coord.row >= len(*mapp) || coord.col < 0 || coord.col >= len((*mapp)[0]) {
				break
			}
			if dir == Up && mapUp[coord.row][coord.col] {
				break
			} else if dir == Down && mapDown[coord.row][coord.col] {
				break
			} else if dir == Left && mapLeft[coord.row][coord.col] {
				break
			} else if dir == Right && mapRight[coord.row][coord.col] {
				break
			}

			if dir == Up {
				mapUp[coord.row][coord.col] = true
			} else if dir == Down {
				mapDown[coord.row][coord.col] = true
			} else if dir == Left {
				mapLeft[coord.row][coord.col] = true
			} else if dir == Right {
				mapRight[coord.row][coord.col] = true
			}

			(*energy)[coord.row][coord.col] = true
			inside := (*mapp)[coord.row][coord.col]
			if inside == '.' {
				coord, dir = calcPointCoordAndDir(coord, dir)
			} else if inside == '|' {
				if dir == Left || dir == Right {
					coord = Coords{coord.row - 1, coord.col}
					dir = Up
					coords = append(coords, Coords{coord.row + 1, coord.col})
					directions = append(directions, Down)
				} else {
					coord, dir = calcPointCoordAndDir(coord, dir)
				}
			} else if inside == '-' {
				if dir == Up || dir == Down {
					coord = Coords{coord.row, coord.col + 1}
					dir = Right
					coords = append(coords, Coords{coord.row, coord.col - 1})
					directions = append(directions, Left)
				} else {
					coord, dir = calcPointCoordAndDir(coord, dir)
				}
			} else if inside == '\\' {
				coord, dir = calcOblRightCoordAndDir(coord, dir)
			} else if inside == '/' {
				coord, dir = calcOblLeftCoordAndDir(coord, dir)
			}
		}
		ind++
	}

}

func calcOblLeftCoordAndDir(coord Coords, dir Direction) (Coords, Direction) {
	if dir == Right {
		return Coords{coord.row - 1, coord.col}, Up
	} else if dir == Left {
		return Coords{coord.row + 1, coord.col}, Down
	} else if dir == Up {
		return Coords{coord.row, coord.col + 1}, Right
	} else {
		return Coords{coord.row, coord.col - 1}, Left
	}
}

func calcOblRightCoordAndDir(coord Coords, dir Direction) (Coords, Direction) {
	if dir == Right {
		return Coords{coord.row + 1, coord.col}, Down
	} else if dir == Left {
		return Coords{coord.row - 1, coord.col}, Up
	} else if dir == Up {
		return Coords{coord.row, coord.col - 1}, Left
	} else {
		return Coords{coord.row, coord.col + 1}, Right
	}
}

func calcPointCoordAndDir(coord Coords, dir Direction) (Coords, Direction) {
	if dir == Right {
		return Coords{coord.row, coord.col + 1}, dir
	} else if dir == Left {
		return Coords{coord.row, coord.col - 1}, dir
	} else if dir == Up {
		return Coords{coord.row - 1, coord.col}, dir
	} else {
		return Coords{coord.row + 1, coord.col}, dir
	}
}

type Coords struct {
	row int
	col int
}

type Direction int64

const (
	Up Direction = iota
	Down
	Right
	Left
)

const inputs = `\..................-...|..............................|./................-.....................\.../..........
.-....../.....-..........|..-.........................|............/.......................\....\......-......
........|...\..............|/.........|...../...............................\../............../...............
...................-.............-.|......-................\........-...............-....\..../\....|..\......
.............|............\....................../..........\..../\...........................................
..............................|.............../..|......-......./........|..............................|..-.\
..............-............/|...\..\./.\............./.............-.......\............./../............-....
.........................................|.........../................................/..........|......|../..
......|.........\..\..............................|.....-\.-.\...........|...............-....-....\..........
...........................\\....|......|..............|....-../.....\|....|\........../..-...................
...........-..../.|.......-|...-.......................|.......-.|..../........................\/.............
.....................-.............\.............../............................................\\............
......-..........-.-..|...........|...........................\...............................|...........|...
................--...\........\.................../................................../............-...........
...............|../............/...../......./..\............|..../-...........|..................\.-.........
|....................|-.\...........\...\.............................................../................\....
\.|................\.\...................\....-./....\......../...................-.....................-.....
...................................................|........../......\.................|..\.............../...
................/.............|......|...........-......................|...-.................../../..|.....-.
.................|-..|.........................-............\....-./......-......../............/.....|.......
.....|.........-.....-.....-................/\..........-........|...\../..|.......-.........-...........-....
......................./.........../........|............\............../........./.....\....-/...............
/........./.......|........|..../../........|............................\........./-..-..............|.....|.
...........|.........-/........-...|/.............................-....|.....\........................|.......
.|..\............/.-..........\........-.............../.......\...........\..............\../-..-|.|-........
..\-|.......................\.......\.\.....\............-...../.......-....../.............................-.
.-........................|........../.....................................-....\.............................
..../.............-.................\.|........................./............/.|..../../................\-....
........../-................./........|......-../........|..-........./........./...........\............\....
......../.......................\..-............-...................\.....................\/........\...../...
....../....|..........-.....-/.................-...\..........................................................
/.........\.............-....|.-..\...../../......./|............................................./...../.....
.......-..|............./.............../|....../..................\....--..-./............./............./..-
-...........|............/.........-..|..................../....................../...........|...\\..........
/...-........../....................................|.......-.....|............|........././..\...............
....................../....\........|.....|....\.......-............|\....\../.....--.....-.|.................
..../........-.....\........\.....\...............|/..................|........./...-......|.................|
.......................-.......//..................|...........|-....-..-...........-.......-....\...../....\.
..................-..../......./-............/.../...\...\../............\\.|.................................
........../.........|...-..........--....................../.......................-......../......\..........
..|........../...........|............./...............|........-....\.....\...............-.....-....../.-...
...../.-..|\...|...............................-...|....\......|.............../.\............\............\..
................\/.....|..........\.................../........\./...................-........................
.........\...../..|............|...............|.-.....................\....\\.|........../.......|...........
......................\\..../.....\.....\..........|....................\.....................................
.\......./............-\........../.........-...../....../....-..............-........................\.....|.
...........\.......-.....\...\/.......-..............-...........\...................../.................\....
............-.../.....|..............-.............|......|.......|..|.............../...|.....\..............
................-..................\..|............./.................\....|..//.......|.../..................
........|.../...|....../...\....................../....-...............-........|..........................|..
..|...../..............................-............|..........-........|.........../..|..|...................
.....-..........-.....-.........................................|.........\....-.....\...........-..../...|...
.-.-........../.................../...........-................|..-............../..-..-......................
/-\.........................-|.........-......|....|-...........-...................................-...\.....
...\.........|.............././...........|./.................-........\.........\.../|........./....\\.......
..|.................-......\./....||.........\....|..........-......|.................\..........\.....-../...
......../.|................................../...................|..\\.....--...................-..|..........
.-........-......./...\....|....../..........-./......./..|................\/......../......../....../........
...-....../.|........./.............\.....|../.-./.......................-....|.........\......|/..........|..
.....|................//.\......|.\/.................................|-...//........./...-.-............\.....
.............../...............\..-....................-..../........../....-...............|-.......\../...\.
..|......\........./../..................|....|................|\.|....\....-.|.......-../......||.-.|....\./.
......|./.................\............|.......................\........-.....-...../..........\......-.../..|
.....|.\....\..............-......-./......./......................../............|..............|............
..............................\..................../................|.................-......./.........../...
.......-.....-..........\..-...../.....\.\..............\..\....../.|.../...............-..................\..
.................|..................-.\....................................\.....-............................
...-.......|........-....-.......-....-..........-.-..\.......................................\.-.../......../
.............-......\................-.....-.....-/......./...\\../.................\....-......./..-.....-|..
....\.............|..-......./-......|............\.........-........\..../...\.......|.|..-...\..............
.....................-....................................../.......\....\........................../\..../\..
........-............\....\./...........\.-........./........./..\.........................../....-..|......|.
..................../.........../.\..............-............\.........\/...\..................-.............
.-...\......\.............................................................-.............\.\.-......|..........
................|..............|.-........../......\.............|............./.|..-|..........|.............
...-........../....\.../.../.........................................................../.........|.-......-...
..-....................-..........................|..../.......|........|../......./....../...-......../....-.
........\............................-.......\...../\......|\......-............./.........-..................
...../...../|..............|.-.....\....-/...........\...............\.\.......................-.-............
......\.......\.....................|................................\.....\..............................\./.
\...............................-.............-../-............................./...............\.|...........
......|......|................-|..-....\/....\...................-...........|...../....................\\....
.............../..............-...|....................\\..............-....\..........\...............|....|.
...\.-.......|..-/...................../|............\......-.../...................../.......................
........../........./.........-....\........................|./......./..............-...-....................
....-|-..............\.......................\.................|.......\.........../..........-......./.......
..-...........................|\.................\................/\................................../.|.....
......\.................................|........\./............-..............\..........-.........../..-....
................-..........|.............|.....|....................-..........\......|-\....|...........-....
...|.........../\............./.........\........\....................|...........................-.\.........
../.|...................././......-..............-\.....-........-.|...........\......-./..-..................
......\.......-..-..............................-................|.................../........................
..\..-................/..-....-........./..........................-.../...........................-..........
............./-./.\........./....-....................././.....|............-...............|./...........-|..
...................-...\.....\....................../............--....................../...........\........
.............-.....\......\./........................../.....................-...................\.........-|.
.....\.........\.../......................\..................|.............................\.|........./.....|
.....\.....\./..............-............../..............\.......|........../..//.-...-.\......../.........\.
..............\...\..........................\...\./.-...|.........|.......|..\../.........................|..
.\..|......\.\\.........../....../....../......./................../..|................/....|.................
./|............./...................................\.-....\..\|./.-.............................\.......\....
\./........./.....\............................/.../...\....................|..............|.|............/|..
..\............\............\.....-..........\.......|../.|......./|......-...--...-..\......|./.-......-.....
...............\........|............\|....|..\.......-................/.........|....\-.|.|..../-.......|....
...................|/.............|............/................|..|.\.|.........|...\........................
..................|../..../.......-................................-......................-.........|.......\.
...............\...../......\./...../............/...........-..........\....../.....\.....\......|...........
../........\.....|..................\......./.......-../.......\.../.|....................................|...
.|.|..\....\.|./.../........................./........-..\.../.............\.............................|....
..................|...\................-..-/..................-...|./..................-..|...................`

const result = 6921
