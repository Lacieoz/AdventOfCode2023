package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(inputs, "\n")

	// MOVE ROCKS UP
	var mapp [][]rune

	// CREATE MAP
	for _, row := range rows {
		mapp = append(mapp, []rune(row))
	}

	// MOVE ROCKS UP
	times := 1000000000
	var previousMaps [][][]rune

	for i := 0; i < times; i++ {
		moveRocksUp(mapp)
		moveRocksLeft(mapp)
		moveRocksDown(mapp)
		moveRocksRight(mapp)

		// FIND POSSIBLE LOOP
		loopLen := findPrevious(previousMaps, mapp)
		if loopLen != 0 {
			ind := len(previousMaps) - ((times - i) % loopLen)
			mapp = previousMaps[ind]
			break
		}
		previousMaps = append(previousMaps, copyMap(mapp))
	}

	res := countRes(mapp)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func findPrevious(previousMaps [][][]rune, mapp [][]rune) int {
	for i := range previousMaps {
		if sameMap(previousMaps[i], mapp) {
			return len(previousMaps) - i
		}
	}
	return 0
}

func sameMap(previous [][]rune, mapp [][]rune) bool {
	for i := range previous {
		for k := range previous[i] {
			if previous[i][k] != mapp[i][k] {
				return false
			}
		}
	}
	return true
}

func copyMap(mapp [][]rune) [][]rune {
	duplicate := make([][]rune, len(mapp))
	for i := range mapp {
		duplicate[i] = make([]rune, len(mapp[i]))
		copy(duplicate[i], mapp[i])
	}
	return duplicate
}

func countRes(mapp [][]rune) interface{} {
	nRows := len(mapp[0])
	res := 0
	for row := range mapp {
		for col := range mapp[row] {
			if mapp[row][col] == 'O' {
				res += nRows - row
			}
		}
	}
	return res
}

func moveRocksUp(mapp [][]rune) {
	for row := range mapp {
		for col := range mapp[row] {
			if mapp[row][col] == 'O' {
				moveRockUp(mapp, row, col)
			}
		}
	}
}

func moveRockUp(mapp [][]rune, row int, col int) {
	for row > 0 {
		if mapp[row-1][col] == '.' {
			mapp[row-1][col] = 'O'
			mapp[row][col] = '.'
		} else {
			break
		}
		row--
	}
}

func moveRocksLeft(mapp [][]rune) {
	for row := range mapp {
		for col := range mapp[row] {
			if mapp[row][col] == 'O' {
				moveRockLeft(mapp, row, col)
			}
		}
	}
}

func moveRockLeft(mapp [][]rune, row int, col int) {
	for col > 0 {
		if mapp[row][col-1] == '.' {
			mapp[row][col-1] = 'O'
			mapp[row][col] = '.'
		} else {
			break
		}
		col--
	}
}

func moveRocksDown(mapp [][]rune) {
	nRows := len(mapp) - 1
	for row := range mapp {
		rowFromDown := nRows - row
		for col := range mapp[row] {
			if mapp[rowFromDown][col] == 'O' {
				moveRockDown(mapp, rowFromDown, col)
			}
		}
	}
}

func moveRockDown(mapp [][]rune, row int, col int) {
	for row < len(mapp)-1 {
		if mapp[row+1][col] == '.' {
			mapp[row+1][col] = 'O'
			mapp[row][col] = '.'
		} else {
			break
		}
		row++
	}
}

func moveRocksRight(mapp [][]rune) {
	for row := range mapp {
		nCols := len(mapp[row]) - 1
		for col := range mapp[row] {
			colFromRight := nCols - col
			if mapp[row][colFromRight] == 'O' {
				moveRockRight(mapp, row, colFromRight)
			}
		}
	}
}

func moveRockRight(mapp [][]rune, row int, col int) {
	for col < len(mapp[row])-1 {
		if mapp[row][col+1] == '.' {
			mapp[row][col+1] = 'O'
			mapp[row][col] = '.'
		} else {
			break
		}
		col++
	}
}

const inputs = `O.........O.......O...OO.O...#.#.O.#O...O.#O..O.##.##.O...#..O..#..O..##O#.#........O#.O..#..#......
.OO..#OO.#.#.OO..#...O.OO.#......#.#OO..O##...##..#O#..#O##O.#....#O#O##.#.#O#.....O.#.O.#....O.....
..O.#.O.O#...#.#....#.#.....OO....#..O#....#.......O..OO..#..#.#O.O.OO...O.O...##...O.#...#.O.....O.
..O.O...#O....O#...O..#OO#....O.......O...#O##.O#....#O.#O.#.#.#...O#.O....O....##O#O#..O..#...#.O.O
##.O#.O..OOO..#....#.#..#..O.O.#.O.OO...##.O.O...#..##.O....O#..##.O.O....OO.....#OOO........O..#.OO
#..OO.#.....OO.........OOO.O..OOO.#O#..#.#O..O#..#OO.O...........O.O.O.......###.....#.#..O...#....#
#.O.O........#....O#..#OO...O..#.#.O..O.......#.O##.O.O........O..O.O.O.O#O.O#.....O.#.....#O.O.#O..
.O..O..O.O.O..#O.....O...O....#O...#...#O...OO..OO...#O..#.##.#O...O..##O...##...O#..#.....#.#...#..
OO..O#..#OOOO...##..#....O.#........#..O...#O......O..O..........#..O..##....OO....O.......#O...O.O.
.#..O.........#......OOO.O...#...O.O..#..OO.#..O..#....##....O.....O#.#O.O.OO...O..#...O....O.......
..#..O.#O.....#...OOO.#.........O.O.#.O...#OO.O...#O..#O..#..#O...#...O......O#......O..#O..O##.....
#.......#O...#O#O...O#O.O.O..............O.........O.O....#O##OO.O#..OO.....#.O..#.O.#.....OOO.#..#.
O..OOOO.......#O.O..O.#..#...O......#..#...#......OO....O..O.O.#....O.O...O..#...##.O.....#.OO..##..
.....#O.....O......#OO.#.O.#..#O.#..............#.O.....#.O.O..OO.#.#.#..O##.O.#O..O...O#...#.......
OOO.OOO#..##........#.#OO..OO#.#..O...O.#.O...##O#.OOO#OO.....#.O..#.O..O......##.#..#.#OO#...#..#..
......#O...O....O......O...#O..O.#...O....O.O...##.O.#....#O..O..............O#..............#.OO.O.
.#....##..O....O.O.#....O...O..O..O..#...#O.O..OOO.....O..#...#........O..#.OOO..O.#..#O.#.O........
O......O...#..O#.........#O..#..........O..#.O.O#.#.####.OO.#.O.....OO.#..#OO.O......##..##.....#.#.
#....##.O...#..##O...O......#OO.#OO.........OO#.O.O....O#....O#..O#...O...O#OO..O....O.#.O#.#..O.O..
......O.#...#O.......#..#...#.OOO..O.##..#.OO..OO...OO#O..O...OO.................#...O....#..OO...#O
..##....#O....O..#O..O..OO.O.O....O#.##..#O..O.O.O..#.OO#..O.....O.O...#OO.....#..#O.#...O#.........
..OO.O.OO.#.O..........#.O.O.#....#.......#.......OOO..#.....#..OO........#..O##O.....#...#O..O.....
.OO#.O#O...O#.O#..#OO.#......#.O.....O..#O...O......O.OO###....OO#.O.........##.#..O.....#..........
.O....O.O..O.....#.#...........#O.O.O#.............O....#....OO....#.....#OO..OO..OOO....O.O..#.OO.#
..O#.O....#..##...#...O....OO#O.#..O.O#..O..O..#..#....#....#..O..##..O.....#...O...#...#....#.OOO#O
.O..O..O...#....O.OO.##..OO#O.#.....#.O..O#O.OO...#.##O..OO.O......##..O.O..#..OO#O#.O#O....O#.....O
.OO...O##......O....##...O..##O##..O..OO##.#...O.##.OO.#.##....#.OO.....O..##.#.#..##O..O..O#..#O.O.
O.#.O...#....O.####OOOO.OO.O.O#..##......O#O....#OO.O.#.....O.....O.##....#....#...O#..O.OOO#O......
##..O.........O.....OO....O#O.O..#..#O..#..OO...OOOO..OO.O.....O##..##.#O............#O....O.O..O...
O.#..O.....#..#.#...O...O#.....O...#.OO...O.....O#.O.....O#O....##....#...O..#O...##..O...O...#..O..
.O##O.#.O..O..#..O...O#....O....O....#..O.....O##OOO.....O###O#.O.O.O.O.#.OO...O......##.#O.....#..O
.#OO.O#.#.#O...OO.OO....#OO.O#O..O..#..O..#.O.O.O..O#O...........O......OOO....#.O......#OO#..#OO#O.
...O#..#....#.O......#..O...OOO........O.#.#.....O.##......#....#...OO....O.O..O....O#.#.O#O.O....##
.....##..OO#..O#.O.#..O....OO.....##.OO..O...O#...###..O#O.O...O.#O#..O.O.....O..O......O......O#..#
O#..O#O...O.OO.OO....O###.#......#O....#O....O#...#O...O.....O#...........OO..O...O......#.#.#.O..#.
........#...#...O.O.#..OO..#O..O..O....OO...#....OO..###...#...#.#....O.OOO..#..O..###.O....O#OO..O.
.#O#O......O.....O#..O.O...O......O#..OO.#.#.O.O.O..O#.#........O.......#..#.O....#.....##........#.
O#O.O...#O.O#O.....#.....#####OOO..O#..#..#.#......O......#....#.....O.O..........#..##..O.........O
..#.O..OO......O.....O.OO..#.....O..##..#....O........#..O.#.......O...#OO#....##..O.#....O.O.#.O..O
.#...O..#OO.#.O...OO.O.##...OOOO....#..O..............O.###.#.O...#.##OOO#O..O....O...O#.###.....O.O
#O.#.O.#.O.....#...O#.O....##.........#..#.....O.....#...O..O.#O.O..O##..OOOO.......O.O#O.......O..O
.#.O.......OO....OO#.OO.#O.O....##.#...##OO....OO...O.###..O.##O..O#...#...#O#...#.#O#.O.#..OO.##..O
#.#.......O...O....#O#.#....#.O.#.....#......O...........#.O..#..O#.##...OO.#...O...#....O.......O..
....OO.O#O#..#O##O..#........#....O...OOO....##.#O..O...#..O.#O##..........O.#.O.....#O.OO.O#.....#.
..O.#O...##.....##..#.OO##.#...OOO.O.O......#.OO....O..#..#O...O.......O.....#O##..#....O#..O.##O.O.
O...O....O....#..O............#O.OO..#..OOO.........OO#.#.O.O..O.OO....O......#.O#.#....O..#O......O
#.......#...OOO.........#.#O.#O....O....O.O#...O#..O.O##O...O.#O..##O....OO........OO#.....#.#.O.#..
.#......O#.O.O..OOO.......O#...#...OO#O.......O.....#O#.#.O..O.OO...OO##.#..O....O.#..##...O.#..#.O#
.#..O......O.OO..#..#O.#..O#.O......#O..O.......##O.#...#....#..#.##O.#O...O....O.....O.#O...O#.O.#.
OOO.#O#..........O#..O.....#...O.O.O..O.O....O...O#....O.O..OO...##.#.O#OO......#...#.O.......O#.O..
....#.#.......O.....##O....O...O....O..O#.O.....O#OO..#OO....#....#...O.......##......O.OO.......OO#
.#........O.O..#.#...O..##.O.##..O.O.#....O..O#.#.#O.......O.......O.#.....O..#..#...#.O....OO.#.O..
.#.O.#.....#....#.#.O.#O..O.OO#..#...OO.......#O#..#O.#......O#.O..#......#.O#.OO...#..O.......O.O..
.#O##O#...O..........OO...OO..OO#...#OO..O....O..OOO...O.....##...O..#..###O.#O.###......O.OO#O#OO..
.##....#........#..O#......OO.....OOO.......O.OO...#O...#.O#..O...OO.OO#.......#.#.#.O#..#O.#.#..O..
#....#.##...O.OO..#.......O..O.O........O...O..OO.......#.O#O#.O.O.OO##.#.#...##.......#....O.O#....
...##....O#.#...O..#OO.O#.O...OO......O#.....#.##..O....#O.#....O..O.....OO#..#.O......#..#.##.#....
O#O......#.#O.........O...OO.#.OO...##.#...OO....O..##.......OOO.O..#...O.O...OO.....#.O..#..#O.....
.O.O#.........#.#...OOOOOO..#....#..O....OO...O..#....#..#.#O.#...O#..#.....O.##.O..#O#......O.O.O..
.#.O#......O##....##.#OO.OO..#...#...O..O....#O..O.#......#..#...#..O.....##...#...O.O.##OOO##......
...OO#O...O....O......O#.##...#.........OOO....O.#O#......OO.#......OOO.#..........O#O....OO.#.O.#.O
....O..O.O...#.......OOO.#...#O.#.....O.#..#.O.#..O.....O#.##..O..#..#...#O.....OO.#O......#....O..O
#....OO.O.O.OO#.O..#...#.#.##.O..OO...O#.OO..O#.#.O.#....O.O......O..##............O............O..O
#O#.O.O.....OO......#..OO...#....O..#.O...#O.#....O.....O.#.O.O.O..#...OO.O......O...#..#..O...#..OO
..#..#.OO#..O.....#....#O#...#......O....#OO.O....#O.....#..#.O....####..#.O..#...OO..#..OO..##..#..
..O...OO....#.#.#....#.O...#.....O.........OO..#.#O.O..O..#O..O.....OO..#.#.#...O....#....OOO.OOO.#O
.....##..O.O.O.O...#..OOO..........O.....##.#..O.#.....O......O.#....#...O....O..O.OOO.#O.....#..###
#...#....O#.O...O.......#.#......#..O.#.........O....OOO.O.O..O..O##..O#O..#..O.#.....###....#.....O
.#..O...#OO...O....O#O....#..###..#.O..##.O...O.....#O.OO......O#....#.....##.#..#......O.OO...##..O
.O....#O#.#...O..O....O#..OO.#O....O..O..#.O#O.O....O####...O..O...#.##...O...#....###OO#.O#..#O#OOO
O.......O#......O#OO...O..##.OO.O....O....#...O.#..OO#....OO..O##.O.OO..#.#O..#.O.#........#...O.#.#
O...O.###O.#..#.#O..O#.##O#.O.O.#..#.....O...#OO..#.#....#....O...#.OOOO.#...#...O.O..O..O....#..OO.
....###.....OO.O...#...O..........#O.##.O.O.OO..O...O#........O.#.#OO#O.........OOO..............#..
O.O.#.O##O..#OOOO.........O.O.O..#..O..O#.O....OO#...OO..#.O....OO.....OO..O..O.O#O#.O###OO..#OO.#.O
OOO..O#O#O#.#O.O#..OO..#......###O....#..O..O....OO....O.###.OOO....O.......O.OO.#...O#O..O#O..#....
..OO......O##..####...#O#O#.O.#..#.....#..OO....#O...........OO...#.O....#O...#....OO.O...#..OO.OO.O
#.....O..#.#..O..........#...#..O....O#........O#..O...O.##.#.O..O...O#O.O....#.....O..OO..#OO..#.#.
O.O...........#.O.O.O.....#.#O#...OO......O.#...........OO..O.#O..#....O.O...#..#OO.O.O#..#..O#OOO#O
.O#.#OO#..#..............OO.#.#........##..O...........O..#...##...OO..#O...#...OO.O.O.#..#..O...OO.
.O.#.........O..........#O.O.......O...##O...OO#O.#O#.#O#..#......O#O..#O.##.O#.#.#..#O.......#..#..
..OOO....O#O.......O....O#....#.....O#...#..O.#.O.O.O.#...#.O.....#.........#.O..#..#.....O..O.O..#.
...OO..O...#........#....O.O#.O#O..O.......#O.......O##...O#.O.#.#O#.....OOO.OOOO.##....O..O..O#.O..
..O...O...O....O..#..OO....OO...#.#...O...#O..O..#.#.....OO.....OOOO......O.##.#..OO#..O#.....O#O.#O
O.O.#O.#..##...O.OO...O.O...#.......#....O##..OO..#.O.O...O.O...##.......OO...#.O###.O...#.....O...O
..OO.....#O#OO##....O.###..##O.##.......O#O.#O..O.O#.O....O.#OO#...#..O.#.....#.O......O#O.O..O##O#.
..............O.O.#...OO.##.OO..OO....#...#..OOO.....#...#............O.O.......OO......OO.#O.#.....
........OO....O.##O.O#..##......O...OO..O#.O...O..........#.........#.....##...#.O#...O.O...........
..#.O...O.....O.OOO.O...##......#.##....#..#..#........#.O.O.O....O.#O.O#........#O..O.......O..O.#.
....OO#.....OO....#...O#..#.......#....O#...O..#..##.O..#..#.....#.#...OO..O..#...#..##.......O.O...
.......#....##O.....OOOO....#.O....OOOO.#.OO.##..O...O#..#.###..#....#......O.O...O.#O#.........O.##
...##O...OO#....#OO...O..#O........#....#...#O...#..#.O...#.........#..O....O..O..O.O..#O...O..#.O#.
...O.O.......OO.......#O.......O...O...OO.O#O.O...O..#..#O......#O.........O.O...O.O....O#O.O.OO..O#
#.#.#.#...#O.OO#..O..#O..#.....#...O.......O.#.O#.#....O.O...#OO.#......O.#..O...O.O...O..#.......#.
....#.#.....O........OO...#......#.O.O...##.OO..O#...#....O#........##...OO.#.....OO...#.#.O..#.#...
....OOO.#......##...O...O#.O..#.O.O...O..#.......OO.....O.#..OO..OOO#..O.#.#.O........OOO.#......#O#
.....#O.O...OO#OOOO#..#...##....O...#.O..O..#...#..OOO.O#O...#..O.#..OO#.#O...O.O.O.O.O.OOO.O..O..#.
...#....#.....O#O...OO.#.#.OOO.O..#...##OO.OO....O.O..O...O...OO.O..O....#O.#..........OO..#.OOO....
.#....#O#O..#O..O#.....#...#O#.....#O.O.#.##..#O.##.......O..#O.O..O........#.O#...#O##...##.......#
...#......O.O...OO#O..#....O....O...#..OOO.OO#..#...OOO.OO...O.#.#.#.O.O..O...#O....#.##....OO#.O..O
#...O.OO#O.O..O...#O.#..#.....O..OO...O..#...#.#.#..O......O..O.O.O.O...#..#O..#..#....#....#....O..`

const result = 95273
