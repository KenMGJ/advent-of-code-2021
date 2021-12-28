package runner

import "fmt"

func (r *Runner) Day25Part1(lines []string) {
	maxHeight, maxWidth, east, south := parseDay25(lines)

	// PrintDay25(maxHeight, maxWidth, east, south)

	move := true
	i := 0

	for move {

		move = false

		newEast := map[string]bool{}
		for i := 0; i < maxHeight; i++ {
			for j := 0; j < maxWidth; j++ {
				key := day25Key(i, j)
				_, ok := east[key]
				if ok {

					newJ := j + 1
					if newJ == maxWidth {
						newJ = 0
					}

					moveKey := day25Key(i, newJ)
					_, o1 := east[moveKey]
					_, o2 := south[moveKey]
					if o1 || o2 {
						newEast[key] = false
					} else {
						newEast[moveKey] = false
						move = true
					}
				}
			}
		}
		east = newEast

		newSouth := map[string]bool{}
		for i := 0; i < maxHeight; i++ {
			for j := 0; j < maxWidth; j++ {
				key := day25Key(i, j)
				_, ok := south[key]
				if ok {

					newI := i + 1
					if newI == maxHeight {
						newI = 0
					}

					moveKey := day25Key(newI, j)
					_, o1 := east[moveKey]
					_, o2 := south[moveKey]
					if o1 || o2 {
						newSouth[key] = false
					} else {
						newSouth[moveKey] = false
						move = true
					}
				}
			}
		}
		south = newSouth

		i++
	}

	// PrintDay25(maxHeight, maxWidth, east, south)
	fmt.Println(i)
}

func PrintDay25(maxHeigth, maxWidth int, east, south map[string]bool) {
	for i := 0; i < maxHeigth; i++ {
		for j := 0; j < maxWidth; j++ {
			key := day25Key(i, j)
			_, o1 := east[key]
			_, o2 := south[key]
			if o1 {
				fmt.Print(">")
			} else if o2 {
				fmt.Print("v")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (r *Runner) Day25Part2(lines []string) {
	fmt.Println("complete remaining days")
}

func parseDay25(lines []string) (int, int, map[string]bool, map[string]bool) {

	maxHeight := len(lines)
	maxWidth := len(lines[0])

	east := map[string]bool{}
	south := map[string]bool{}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			switch string(lines[i][j]) {
			case ">":
				east[day25Key(i, j)] = false
			case "v":
				south[day25Key(i, j)] = false
			}
		}
	}

	return maxHeight, maxWidth, east, south
}

func day25Key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
