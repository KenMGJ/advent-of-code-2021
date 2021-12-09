package runner

import (
	"fmt"
	"strconv"
	"strings"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInteger = 1<<(UintSize-1) - 1

func (r *Runner) Day09Part1(lines []string) {
	points := inputToDay09PointArray(lines)
	sum := 0
	for _, p := range points {
		for _, lp := range p {
			if lp.IsLowPoint() {
				// fmt.Println(lp)
				sum += lp.RiskValue()
			}
		}
	}
	fmt.Println(sum)
}

var foundDay09Map map[*Day09Point]bool

func (r *Runner) Day09Part2(lines []string) {
	points := inputToDay09PointArray(lines)

	basins := []int{}

	foundDay09Map = make(map[*Day09Point]bool)
	height := len(points)
	for i := 0; i < height; i++ {
		length := len(points[i])
		for j := 0; j < length; j++ {
			// fmt.Printf("i = %d ; j = %d\n", i, j)

			size := basinSize(points, height, length, i, j)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}

	m1, m2, m3 := 0, 0, 0
	for _, b := range basins {
		if b > m1 {
			m3 = m2
			m2 = m1
			m1 = b
		} else if b > m2 {
			m3 = m2
			m2 = b
		} else if b > m3 {
			m3 = b
		}
	}

	fmt.Println(m1 * m2 * m3)
}

func basinSize(points [][]*Day09Point, height, length, i, j int) int {
	point := points[i][j]
	_, ok := foundDay09Map[point]
	if ok {
		return 0
	}

	foundDay09Map[point] = true
	if point.Height == 9 {
		return 0
	}

	// fmt.Printf("\ti = %d ; j = %d\n", i, j)

	basinSizeUp := 0
	if i > 0 {
		basinSizeUp += basinSize(points, height, length, i-1, j)
	}

	basinSizeDown := 0
	if i < height-1 {
		basinSizeDown += basinSize(points, height, length, i+1, j)
	}

	basinSizeLeft := 0
	if j > 0 {
		basinSizeLeft += basinSize(points, height, length, i, j-1)
	}

	basinSizeRight := 0
	if j < length-1 {
		basinSizeRight += basinSize(points, height, length, i, j+1)
	}

	return 1 + basinSizeUp + basinSizeDown + basinSizeLeft + basinSizeRight
}

type Day09Point struct {
	Height int
	Up     int
	Down   int
	Left   int
	Right  int
}

func (d *Day09Point) IsLowPoint() bool {
	return d.Height < d.Up && d.Height < d.Down && d.Height < d.Left && d.Height < d.Right
}

func (d *Day09Point) RiskValue() int {
	return 1 + d.Height
}

func inputToDay09PointArray(lines []string) [][]*Day09Point {
	points := [][]*Day09Point{}

	ltos := [][]string{}
	for _, l := range lines {
		ltos = append(ltos, strings.Split(l, ""))
	}

	for i := 0; i < len(ltos); i++ {
		linePoints := []*Day09Point{}
		for j := 0; j < len(ltos[i]); j++ {

			up := MaxInteger
			if i > 0 {
				u, err := strconv.Atoi(ltos[i-1][j])
				if err != nil {
					panic(err)
				}
				up = u
			}

			down := MaxInteger
			if i < len(ltos)-1 {
				d, err := strconv.Atoi(ltos[i+1][j])
				if err != nil {
					panic(err)
				}
				down = d
			}

			left := MaxInteger
			if j > 0 {
				l, err := strconv.Atoi(ltos[i][j-1])
				if err != nil {
					panic(err)
				}
				left = l
			}

			right := MaxInteger
			if j < len(ltos[i])-1 {
				r, err := strconv.Atoi(ltos[i][j+1])
				if err != nil {
					panic(err)
				}
				right = r
			}

			height, err := strconv.Atoi(ltos[i][j])
			if err != nil {
				panic(err)
			}

			linePoints = append(linePoints, &Day09Point{Height: height, Up: up, Down: down, Left: left, Right: right})
		}
		points = append(points, linePoints)
	}

	return points
}
