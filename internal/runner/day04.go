package runner

import (
	"fmt"
	"strconv"
	"strings"
)

func (r *Runner) Day04Part1(lines []string) {
	draws, boards := parseDay04Input(lines)

	for _, draw := range draws {
		markDraw(boards, draw)

		bingo := false
		// determine if bingo
		for _, b := range boards {
			if isBingo(b) {
				bingo = true
				sum := sumUnmarked(b)
				fmt.Println(sum * draw)
				break
			}
		}

		if bingo {
			break
		}
	}
}

func markDraw(boards [][][]*bingoTile, draw int) {
	for _, b := range boards {
		for _, r := range b {
			for _, c := range r {
				if c.value == draw {
					c.marked = true
				}
			}
		}
	}
}

func (r *Runner) Day04Part2(lines []string) {
	draws, boards := parseDay04Input(lines)

	for _, draw := range draws {
		markDraw(boards, draw)

		bingo := false
		for i := 0; i < len(boards); i++ {
			board := boards[i]
			if isBingo(board) {
				if len(boards) == 1 {
					bingo = true
					sum := sumUnmarked(board)
					fmt.Println(sum * draw)
					break
				} else {
					copy(boards[i:], boards[i+1:])
					boards[len(boards)-1] = [][]*bingoTile{}
					boards = boards[:len(boards)-1]
				}
			}
		}

		if bingo {
			break
		}
	}
}

type bingoTile struct {
	value  int
	marked bool
}

func parseDay04Input(lines []string) ([]int, [][][]*bingoTile) {

	// number draws
	draws := []int{}
	drawLine := lines[0]
	for _, s := range strings.Split(drawLine, ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		draws = append(draws, v)
	}

	lines = lines[1:]

	boards := [][][]*bingoTile{}
	var board [][]*bingoTile

	boardNum := 0
	boardRow := 0

	for _, l := range lines {
		if l == "" {
			if board != nil {
				boards = append(boards, board)
			}
			board = [][]*bingoTile{}
			boardNum++
			boardRow = 0
			continue
		}

		vals := strings.Split(l, " ")
		board = append(board, []*bingoTile{})

		for _, s := range vals {
			if s == "" || s == " " {
				continue
			}

			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			space := bingoTile{
				value:  val,
				marked: false,
			}

			board[boardRow] = append(board[boardRow], &space)
		}

		boardRow++
	}

	boards = append(boards, board)
	return draws, boards
}

func printBingoBoards(boards [][][]*bingoTile) {
	for i, board := range boards {
		fmt.Printf("Board #%d\n", i+1)
		for _, row := range board {
			for _, col := range row {
				s := "X"
				if !col.marked {
					s = fmt.Sprintf("%2d", col.value)
				}
				fmt.Printf("%s ", s)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

const bingoBoardRows = 5
const bingoBoardCols = 5

func isBingo(board [][]*bingoTile) bool {
	for i := 0; i < bingoBoardRows; i++ {
		found := true
		for j := 0; j < bingoBoardCols; j++ {
			if !board[i][j].marked {
				found = false
				continue
			}
		}
		if found {
			return true
		}
	}

	for j := 0; j < bingoBoardCols; j++ {
		found := true
		for i := 0; i < bingoBoardRows; i++ {
			if !board[i][j].marked {
				found = false
				continue
			}
		}
		if found {
			return true
		}
	}

	return false
}

func sumUnmarked(board [][]*bingoTile) int {
	sum := 0
	for _, row := range board {
		for _, col := range row {
			if !col.marked {
				sum += col.value
			}
		}
	}
	return sum
}
