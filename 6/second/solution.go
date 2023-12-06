package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(inputs, "\n")
	var timesStr = strings.Split(rows[0], ":")
	timesStrNoSpaces := strings.Replace(timesStr[1], " ", "", -1)
	var distancesStr = strings.Split(rows[1], ":")
	distancesStrNoSpaces := strings.Replace(distancesStr[1], " ", "", -1)
	t, _ := strconv.Atoi(timesStrNoSpaces)
	d, _ := strconv.Atoi(distancesStrNoSpaces)

	var res = disequationSolution(t, d)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func disequationSolution(t int, d int) int {
	radix1, radix2 := secondGradeEq(t, d)
	return radix2 - radix1 - 1
}

func secondGradeEq(t int, d int) (int, int) {
	deltaSqrt := math.Sqrt(float64((t * t) - (4 * d)))
	res1 := (float64(t) - deltaSqrt) / 2
	res2 := (float64(t) + deltaSqrt) / 2
	return int(math.Floor(res1)), int(math.Ceil(res2))
}

func bruteForceSolution(t int, d int) int {
	var res = 0
	for t1 := 0; t1 < t; t1++ {
		if (t-t1)*t1 > d {
			res++
		}
	}
	return res
}

const inputs = `Time:        44     70     70     80
Distance:   283   1134   1134   1491`

const result = 29432455
