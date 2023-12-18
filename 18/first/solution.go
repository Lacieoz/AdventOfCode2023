package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(inputs, "\n")

	var currCoord = Coords{0, 0}
	var digOps []DigOp
	var allCoords []Coords
	allCoords = append(allCoords, Coords{0, 0})
	var allDirections []Direction
	allDirections = append(allDirections, Right)

	for _, row := range rows {
		digOp := createDigOp(row)
		digOps = append(digOps, digOp)
		for i := digOp.l; i > 0; i-- {
			currCoord = calcCoords(currCoord, digOp)
			allCoords = append(allCoords, currCoord)
			allDirections = append(allDirections, digOp.dir)
		}
	}

	minRow := 0
	maxRow := 0
	minCol := 0
	maxCol := 0

	for _, coord := range allCoords {
		if coord.row < minRow {
			minRow = coord.row
		} else if coord.row > maxRow {
			maxRow = coord.row
		}

		if coord.col < minCol {
			minCol = coord.col
		} else if coord.col > maxCol {
			maxCol = coord.col
		}
	}

	// CREATE MAP
	var mask [][]rune
	for i := 0; i <= (maxRow-minRow)+2; i++ {
		var newRow []rune
		for k := 0; k <= (maxCol-minCol)+2; k++ {
			newRow = append(newRow, '.')
		}
		mask = append(mask, newRow)
	}
	for i, coord := range allCoords {
		row := (coord.row - minRow) + 1
		col := (coord.col - minCol) + 1
		mask[row][col] = '#'
		if allDirections[i] == Up {
			if mask[row][col-1] != '#' {
				mask[row][col-1] = 'A'
			}
			if mask[row][col+1] != '#' {
				mask[row][col+1] = 'B'
			}
		} else if allDirections[i] == Down {
			if mask[row][col-1] != '#' {
				mask[row][col-1] = 'B'
			}
			if mask[row][col+1] != '#' {
				mask[row][col+1] = 'A'
			}
		} else if allDirections[i] == Right {
			if mask[row-1][col] != '#' {
				mask[row-1][col] = 'A'
			}
			if mask[row+1][col] != '#' {
				mask[row+1][col] = 'B'
			}
		} else {
			if mask[row-1][col] != '#' {
				mask[row-1][col] = 'B'
			}
			if mask[row+1][col] != '#' {
				mask[row+1][col] = 'A'
			}
		}
	}

	var external rune
	for i := 0; i < len(mask[0]); i++ {
		if mask[0][i] == 'A' || mask[0][i] == 'B' {
			external = mask[0][i]
			break
		}
	}

	expandExternals(mask, external)

	// CALCULATE TOTAL AREA
	res := 0
	for i := range mask {
		for k := range mask[0] {
			if mask[i][k] == external {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}

			if mask[i][k] != external {
				res++
			}
		}
		fmt.Println("")
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func expandExternals(mask [][]rune, external rune) {
	for i := range mask {
		for k := range mask[0] {
			if mask[i][k] == external {
				expandExternalValue(mask, i, k)
			}
		}
	}
}

func expandExternalValue(mask [][]rune, row int, col int) {
	value := mask[row][col]
	if row-1 >= 0 && mask[row-1][col] != '#' && mask[row-1][col] != value {
		mask[row-1][col] = value
		expandExternalValue(mask, row-1, col)
	}
	if row+1 < len(mask) && mask[row+1][col] != '#' && mask[row+1][col] != value {
		mask[row+1][col] = value
		expandExternalValue(mask, row+1, col)
	}
	if col-1 >= 0 && mask[row][col-1] != '#' && mask[row][col-1] != value {
		mask[row][col-1] = value
		expandExternalValue(mask, row, col-1)
	}
	if col+1 < len(mask[0]) && mask[row][col+1] != '#' && mask[row][col+1] != value {
		mask[row][col+1] = value
		expandExternalValue(mask, row, col+1)
	}
}

func calcCoords(coord Coords, digOp DigOp) Coords {
	digOp.l--
	if digOp.l < 0 {
		fmt.Println("digOp too low", digOp.l)
		syscall.Exit(1)
	}
	if digOp.dir == Up {
		return Coords{coord.row - 1, coord.col}
	} else if digOp.dir == Down {
		return Coords{coord.row + 1, coord.col}
	} else if digOp.dir == Right {
		return Coords{coord.row, coord.col + 1}
	} else {
		return Coords{coord.row, coord.col - 1}
	}
}

func createDigOp(row string) DigOp {
	split := strings.Split(row, " ")
	length, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("ERROR", err)
	}
	return DigOp{length, createDirection(split[0])}
}

func createDirection(input string) Direction {
	if input == "U" {
		return Up
	} else if input == "D" {
		return Down
	} else if input == "R" {
		return Right
	} else {
		return Left
	}
}

type DigOp struct {
	l   int
	dir Direction
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

const inputs = `R 6 (#6248a0)
U 2 (#46d3f3)
R 3 (#7fccf2)
U 6 (#3e2833)
L 2 (#022ef2)
U 3 (#1f9513)
L 11 (#1adba2)
U 2 (#560cf3)
L 12 (#61a9d2)
U 5 (#23fd43)
R 8 (#455242)
U 10 (#720cc3)
L 3 (#7820f2)
U 2 (#720cc1)
L 9 (#47d912)
U 2 (#8308d3)
L 5 (#5bb362)
U 4 (#95d253)
R 3 (#0c4320)
U 6 (#5e4863)
R 4 (#2918b0)
U 3 (#5ba951)
R 10 (#4aee60)
U 4 (#5ba953)
L 8 (#31be50)
U 6 (#5ef1d3)
L 4 (#857070)
U 4 (#10e113)
R 6 (#65be60)
U 9 (#037461)
L 6 (#27d6e0)
U 3 (#8c6ab1)
L 6 (#27d6e2)
D 7 (#26ac21)
R 5 (#2cdad0)
D 2 (#566133)
L 5 (#151030)
D 7 (#1ae053)
L 6 (#151032)
U 5 (#4549b3)
L 3 (#463a90)
U 12 (#322e43)
L 2 (#4f3440)
U 5 (#403703)
R 12 (#026f32)
U 5 (#4cb003)
R 5 (#026f30)
U 9 (#50cc93)
R 4 (#6248a2)
U 4 (#6994a3)
L 2 (#145ac2)
U 5 (#723e43)
L 9 (#0dcb32)
D 6 (#4f20e3)
L 7 (#61ea42)
D 5 (#3fe981)
L 8 (#2f42c0)
U 3 (#3aafa1)
L 8 (#2f42c2)
U 8 (#46c601)
L 8 (#4e7ef2)
U 4 (#0275c3)
L 2 (#43ea02)
U 5 (#578d43)
R 3 (#2f9202)
U 4 (#3e3dc3)
R 5 (#450192)
U 7 (#3e3dc1)
R 9 (#4e11f2)
U 5 (#563023)
R 9 (#268932)
U 6 (#36c551)
L 9 (#0d5852)
U 5 (#623eb1)
R 5 (#0d5850)
U 7 (#172f21)
R 9 (#26ec72)
D 3 (#10fc93)
R 3 (#7e19d2)
D 7 (#5370f3)
R 4 (#1878b0)
D 6 (#3c0843)
L 7 (#6d8630)
D 6 (#48cf03)
R 6 (#106082)
D 8 (#5f4fa3)
R 7 (#307e12)
U 10 (#012723)
R 4 (#41c6c2)
U 5 (#8dcf43)
R 11 (#12b352)
U 7 (#003e83)
R 5 (#2f6252)
U 11 (#227633)
R 5 (#2c9470)
U 7 (#58e3f3)
R 4 (#982680)
U 5 (#4ea653)
R 3 (#6b7310)
U 3 (#21ca43)
R 3 (#424120)
U 2 (#0a5fc3)
R 8 (#7ab340)
U 7 (#003451)
R 3 (#18f1d0)
U 9 (#9191d1)
R 3 (#066440)
U 6 (#0a8c11)
R 7 (#96e700)
U 3 (#39e521)
R 3 (#1e5b60)
D 7 (#328fe3)
R 6 (#227080)
D 2 (#1f38b3)
R 8 (#8f45a0)
D 6 (#5bb0f3)
R 3 (#1a0f32)
U 6 (#5c2df3)
R 2 (#62c292)
U 6 (#5c2df1)
R 3 (#34e462)
D 5 (#28bdd3)
R 5 (#03b860)
D 7 (#083ef3)
R 3 (#1df6f0)
D 5 (#38cba3)
R 2 (#3dce10)
D 10 (#30bfe3)
R 7 (#3dce12)
D 8 (#595993)
R 9 (#487d20)
D 5 (#6a9523)
R 4 (#479ca0)
D 12 (#24dcf1)
R 6 (#5b1660)
D 7 (#24dcf3)
L 8 (#134720)
D 7 (#051c93)
L 11 (#2a2ab2)
D 7 (#1398c3)
R 7 (#4a1c02)
D 8 (#1398c1)
R 5 (#2ef222)
U 8 (#2e2f53)
R 7 (#1030a2)
D 3 (#1c61f3)
R 8 (#5dbc60)
D 5 (#6f2cc3)
R 4 (#5dbc62)
D 5 (#24a243)
R 3 (#20ece2)
D 3 (#06e543)
R 7 (#a016e2)
D 3 (#1d4f73)
R 6 (#33a030)
U 5 (#2a7df3)
R 2 (#21d750)
U 6 (#5bdf73)
R 9 (#21d752)
U 8 (#1dd4a3)
R 3 (#33a032)
D 4 (#2ddb83)
R 4 (#23edd0)
D 10 (#0c73a3)
R 3 (#14a320)
U 9 (#54ba73)
R 3 (#75d9b0)
U 5 (#30cd13)
R 3 (#2945e0)
U 8 (#5c0323)
R 3 (#2555c0)
U 2 (#6080b3)
R 3 (#5a4420)
U 7 (#10bbe3)
R 4 (#1d22d0)
U 5 (#84e443)
R 7 (#4e4e10)
U 3 (#189573)
R 5 (#65c630)
D 5 (#393a73)
R 4 (#405660)
D 9 (#40e7d3)
R 5 (#5bb560)
D 11 (#7a2241)
R 3 (#38aea0)
U 4 (#8facb3)
R 4 (#4fc0e0)
U 6 (#594c91)
L 4 (#836510)
U 5 (#61c931)
R 5 (#836512)
U 5 (#224851)
R 6 (#11f220)
D 6 (#88b791)
R 7 (#6b9330)
D 5 (#1701c1)
L 7 (#479120)
D 9 (#5c9b93)
R 7 (#127902)
D 8 (#534473)
R 8 (#127900)
U 3 (#3ed6a3)
R 3 (#350ab0)
U 6 (#205a83)
L 4 (#51c5f0)
U 9 (#69d333)
R 4 (#4aa070)
U 9 (#043313)
R 5 (#17f710)
D 4 (#7c5d91)
R 9 (#298c30)
D 5 (#0a8311)
R 3 (#591520)
D 3 (#0a8313)
R 4 (#212a70)
D 3 (#2a5411)
L 4 (#5413a0)
D 2 (#6168e1)
L 8 (#4435d0)
D 5 (#1a8fa1)
L 4 (#441700)
D 5 (#606ed1)
R 11 (#45e2e0)
D 3 (#398c91)
R 2 (#31b3e0)
D 3 (#0fea51)
L 11 (#3c10c0)
D 3 (#3fbb81)
L 3 (#6742e0)
D 5 (#711f21)
L 9 (#49c5e0)
D 3 (#3f2e53)
L 3 (#312ec0)
D 3 (#6d1533)
L 4 (#48b6e0)
D 11 (#296c31)
L 6 (#6d05a0)
D 5 (#47c8e1)
L 2 (#2508c0)
D 5 (#1bf8a1)
L 6 (#920e62)
D 4 (#1f15d1)
L 4 (#373750)
D 10 (#1afa51)
L 3 (#52aa32)
U 14 (#05b921)
L 3 (#6031c2)
D 6 (#05b923)
L 5 (#4806e2)
D 6 (#6fd5f1)
L 6 (#6742e2)
D 10 (#5e1771)
L 3 (#7123e0)
U 5 (#0415b1)
L 3 (#37fe00)
U 8 (#057521)
L 3 (#2e54f0)
U 4 (#5432d1)
L 5 (#2e54f2)
U 6 (#4c2f21)
L 6 (#45bea0)
U 4 (#0c2771)
L 9 (#0d0650)
U 4 (#49b9e1)
L 4 (#8a2940)
U 4 (#09a231)
L 6 (#30a7e0)
U 4 (#631ea1)
L 4 (#637be0)
D 10 (#6cc0d3)
L 4 (#397f50)
D 3 (#8b2cd1)
L 3 (#0dc4e2)
U 8 (#037b61)
L 3 (#2a3d52)
U 5 (#0692c1)
L 5 (#5663d2)
D 3 (#7849d1)
L 2 (#5608c2)
D 12 (#0ef5a3)
L 4 (#161a50)
U 7 (#90e6b3)
L 11 (#161a52)
U 5 (#36cd03)
R 11 (#6c3392)
U 7 (#2170a1)
L 8 (#0a1780)
U 8 (#4a27b1)
L 10 (#0a1782)
D 7 (#6b1101)
L 4 (#3183a2)
D 4 (#21f7c1)
L 9 (#560490)
D 6 (#529ee1)
R 4 (#48a350)
D 3 (#529ee3)
R 9 (#551810)
D 7 (#76b501)
L 9 (#5663d0)
D 4 (#7b0441)
L 3 (#175852)
D 4 (#39bd01)
R 6 (#6deeb2)
D 7 (#2f6b11)
R 3 (#366222)
D 7 (#692813)
R 9 (#0cf6d2)
U 7 (#710a11)
R 7 (#04ccf2)
U 6 (#176a11)
R 3 (#7dff50)
D 6 (#603861)
R 6 (#7dff52)
D 4 (#5b6921)
L 4 (#633e40)
D 5 (#6ca0a1)
L 7 (#573220)
D 6 (#540841)
L 4 (#800082)
D 4 (#3ef661)
L 6 (#0eece2)
D 7 (#35e4f1)
L 7 (#8eed60)
D 3 (#203c11)
R 4 (#573222)
D 9 (#0d67e1)
R 8 (#633e42)
D 4 (#12e141)
R 5 (#5b19c2)
D 6 (#478941)
R 4 (#49f642)
D 8 (#48dfe1)
R 8 (#308d22)
D 2 (#22a041)
R 5 (#1d1560)
D 7 (#a13221)
R 9 (#0d6e40)
D 4 (#7c87c3)
R 10 (#41d430)
D 3 (#218853)
R 2 (#24aa70)
D 10 (#390d31)
R 5 (#3b10b0)
D 7 (#390d33)
L 11 (#4fd880)
D 3 (#228c91)
L 3 (#54cf00)
U 10 (#7b8381)
L 3 (#5c8ce0)
U 4 (#4564c1)
L 3 (#31ef32)
D 8 (#387531)
L 2 (#1d8930)
D 6 (#71a371)
L 4 (#1d8932)
D 5 (#6a5831)
R 3 (#751a72)
D 6 (#031151)
R 6 (#1d68b2)
U 6 (#9dd363)
R 7 (#0cb252)
D 6 (#5978a3)
R 5 (#576582)
D 9 (#203623)
R 3 (#381cd2)
D 3 (#27d351)
R 7 (#26e4a2)
D 3 (#7eee11)
R 9 (#465982)
D 7 (#57f6d3)
R 8 (#4d7762)
D 7 (#3d85d3)
L 7 (#3936e0)
D 3 (#991e33)
L 4 (#3936e2)
U 3 (#228e43)
L 4 (#2bec72)
U 4 (#455891)
L 9 (#4970f2)
D 7 (#73b611)
L 3 (#50b0d2)
D 3 (#981a71)
R 12 (#54b282)
D 6 (#7eee13)
R 8 (#2e1522)
D 3 (#411a11)
R 2 (#02d840)
D 6 (#195871)
R 6 (#02d842)
D 9 (#478381)
R 4 (#4fc650)
D 7 (#9bff91)
R 6 (#2a5240)
U 9 (#1b76c1)
R 3 (#4bdce0)
D 9 (#286ff1)
R 8 (#130cf0)
D 8 (#687c93)
L 5 (#392f80)
D 5 (#7769b3)
L 7 (#45f640)
U 5 (#4b9f21)
L 5 (#8c0a30)
D 3 (#0e98e1)
R 2 (#4053a0)
D 6 (#0c5ae1)
R 8 (#27c960)
D 2 (#4b72f1)
R 5 (#51b6e0)
D 4 (#4c8e51)
R 8 (#100d40)
D 3 (#005121)
R 6 (#8355f0)
D 3 (#0e7391)
L 15 (#32c8e0)
D 3 (#0027c1)
R 9 (#1697b2)
D 7 (#1cd631)
R 5 (#844ef2)
D 5 (#5622f1)
R 6 (#9ae6a0)
D 3 (#4dd421)
L 8 (#4055d2)
D 5 (#3873b1)
L 8 (#014c20)
D 4 (#490ad1)
L 4 (#8f5cf0)
D 4 (#269f71)
L 5 (#0561e0)
D 3 (#6ead01)
R 11 (#69fbf0)
D 7 (#31e0c3)
L 11 (#4bff70)
D 6 (#31e0c1)
L 4 (#32c140)
D 3 (#6ead03)
L 9 (#4468d0)
U 5 (#389cf1)
R 5 (#253bf2)
U 5 (#6b59c1)
R 5 (#6ae392)
U 3 (#5be451)
L 9 (#527ba2)
U 4 (#779f31)
L 4 (#157f92)
U 5 (#0fc333)
R 3 (#44e432)
U 2 (#28c343)
R 10 (#2e1922)
U 3 (#464983)
L 5 (#2e1920)
U 4 (#44b003)
L 5 (#67e3c2)
U 3 (#594071)
L 9 (#14ac92)
U 8 (#29a541)
L 3 (#14ac90)
D 11 (#409a41)
L 2 (#1e4dc2)
D 3 (#1b0aa1)
L 11 (#4fc682)
D 6 (#2dc591)
R 3 (#710c12)
D 8 (#5bfcf1)
R 10 (#3d4570)
D 6 (#006fc3)
L 2 (#2490c0)
D 2 (#45ab51)
L 11 (#657f00)
D 6 (#45ab53)
L 6 (#25b9c0)
U 4 (#006fc1)
L 5 (#355130)
U 3 (#471781)
L 4 (#06ef70)
U 5 (#9b20e1)
L 5 (#3f8cb2)
U 9 (#6cd361)
L 6 (#198892)
D 5 (#47d881)
L 5 (#597ef2)
D 12 (#4c8323)
L 5 (#5fd782)
D 6 (#4c8321)
L 2 (#00bba2)
D 9 (#47d883)
L 4 (#162842)
D 7 (#658221)
L 4 (#5f7a30)
D 7 (#04b5b1)
L 4 (#41c9e0)
D 10 (#69cb31)
L 8 (#1f8e80)
U 6 (#10eca1)
L 2 (#25fd12)
U 5 (#430193)
L 10 (#403f42)
U 2 (#36b173)
L 11 (#5723f2)
U 6 (#79b301)
L 6 (#0e61e2)
D 5 (#72e903)
L 10 (#006a82)
U 5 (#3277a3)
L 3 (#1898f2)
U 8 (#5ffeb3)
R 8 (#951c52)
U 4 (#1ac9b3)
R 4 (#2434e2)
U 4 (#08c763)
R 3 (#6b8ac2)
U 5 (#08c761)
R 4 (#4b9ee2)
D 2 (#2ed673)
R 9 (#26ce02)
D 7 (#5b8f13)
R 3 (#396c32)
D 4 (#5a2f03)
R 4 (#6ca852)
U 12 (#39a223)
R 5 (#037f72)
U 3 (#93d121)
R 13 (#3ef872)
U 4 (#426763)
R 2 (#1d76c2)
U 4 (#141b71)
L 7 (#423682)
U 2 (#4ba9d1)
L 8 (#14f692)
U 4 (#4ba9d3)
L 5 (#4ee6d2)
U 8 (#141b73)
R 7 (#3233b2)
U 9 (#3d83a3)
L 7 (#2ea882)
U 8 (#8534c1)
L 9 (#56c182)
U 5 (#8534c3)
R 10 (#5221d2)
U 7 (#22e323)
R 3 (#27b6f2)
D 5 (#614be3)
R 9 (#09e122)
U 5 (#10a9e3)
R 4 (#6a69e2)
U 4 (#5a1bf3)
R 5 (#20e032)
D 7 (#0cebd3)
R 5 (#15b5c2)
D 4 (#339fc3)
R 7 (#185b22)
D 8 (#2e6693)
R 7 (#39dd10)
D 4 (#438fd3)
R 4 (#895f00)
D 5 (#29bc73)
R 3 (#37b2b2)
U 3 (#19cce3)
R 4 (#551f52)
U 7 (#76b423)
R 7 (#551f50)
U 3 (#351c23)
L 11 (#58c762)
U 4 (#085781)
R 6 (#19bf92)
U 3 (#6dca51)
R 5 (#19bf90)
U 11 (#4f7b51)
L 6 (#3d89c2)
U 6 (#73b973)
L 8 (#0cef60)
U 10 (#4e3903)
L 3 (#23ec20)
U 5 (#0de323)
L 6 (#7ad8d0)
U 4 (#0de321)
L 6 (#8975a0)
D 4 (#4e3901)
L 6 (#5990f0)
D 10 (#2ebf13)
L 6 (#3e8bb0)
D 6 (#603b73)
L 3 (#5c1812)
U 2 (#2b55e3)
L 6 (#468b92)
U 7 (#2b0cd3)
L 4 (#271260)
U 4 (#510f03)
L 5 (#587ac0)
U 7 (#2e54d3)
L 8 (#4f4330)
U 8 (#15dae3)
L 4 (#0b1cf2)
U 9 (#081bb3)
L 3 (#1d8872)
U 3 (#548fe3)
L 5 (#768472)
U 8 (#13ceb3)
L 4 (#2f07f2)
U 8 (#093013)
L 5 (#009e92)
U 6 (#868ba3)
L 5 (#505932)
U 4 (#44e113)
L 9 (#1302c2)
U 5 (#7cc513)
L 9 (#5710b2)
U 8 (#100593)
L 3 (#22c402)
U 4 (#0c32c1)
L 5 (#43f6a2)
U 7 (#9dfcd1)
L 3 (#393c92)
U 3 (#277c21)
R 9 (#1f1802)
U 8 (#024931)
R 3 (#23afa2)
D 3 (#5259b1)
R 4 (#45a2e0)
D 11 (#5114c1)
R 4 (#347180)
U 8 (#9d0ea1)
R 4 (#347182)
U 6 (#50c621)
R 4 (#45a2e2)
D 8 (#1553a1)
R 7 (#23afa0)
U 3 (#17a171)
L 4 (#2de572)
U 8 (#2b55e1)
L 5 (#25dd42)
U 2 (#31fac3)
L 5 (#4ef822)
U 3 (#7b9e93)
R 10 (#3581a2)
U 3 (#01c8c3)
R 15 (#8479c0)
U 4 (#079a33)
L 4 (#3f1bf2)
U 8 (#4ada83)
L 11 (#049b82)
U 5 (#524563)
L 4 (#9bffc2)
U 6 (#313523)
L 6 (#0bf302)
U 3 (#2361b1)
L 14 (#0c52b2)
D 5 (#574911)
R 8 (#94fd32)
D 4 (#0e2ea1)
R 9 (#7681e0)
D 6 (#176361)
L 9 (#510690)
D 8 (#3319b1)
L 8 (#44e302)
D 9 (#3b9491)
L 7 (#44e300)
U 8 (#3ae6e1)
L 2 (#510692)
U 3 (#2d1f31)
L 8 (#7681e2)
U 2 (#0e6031)
L 4 (#67bc32)
U 9 (#8a5723)`

const result = 45159
