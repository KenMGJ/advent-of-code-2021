package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day21Part1(lines []string) {
	startingPositions := parseDay21(lines)

	die := NewDay21DeterministicDice()
	space := NewDay21Space(10)

	pos := map[int]*Day21Space{}
	score := map[int]int{}

	pos[1] = space.Advance(startingPositions[1] - 1)
	score[1] = 0

	pos[2] = space.Advance(startingPositions[2] - 1)
	score[2] = 0

	scoreLimit := 1000

	turn := 0
	player := 0
	for score[1] < scoreLimit && score[2] < scoreLimit {

		if turn%2 == 0 {
			player = 1
		} else {
			player = 2
		}

		move := 0
		move += die.Roll()
		move += die.Roll()
		move += die.Roll()

		pos[player] = pos[player].Advance(move)
		score[player] += pos[player].Value

		turn++
	}

	loser := 1
	if player == 1 {
		loser = 2
	}

	fmt.Println(score[loser] * die.RollCount)
}

func (r *Runner) Day21Part2(lines []string) {
	startingPositions := parseDay21(lines)

	gameCache = map[string]GameResult{}
	results := Day21Part2Run(startingPositions[1], 0, startingPositions[2], 0, 0, 0)

	fmt.Println(util.MaxInt(results.P1Wins, results.P2Wins))
}

type Day21Space struct {
	Value int
	prev  *Day21Space
	next  *Day21Space
}

func (d *Day21Space) Advance(n int) *Day21Space {
	current := d
	for i := 0; i < n; i++ {
		current = current.next
	}
	return current
}

func NewDay21Space(total int) *Day21Space {
	one := &Day21Space{
		Value: 1,
	}

	current := one
	for i := 2; i <= total; i++ {
		next := &Day21Space{
			Value: i,
			prev:  current,
		}
		current.next = next
		current = next
	}

	current.next = one
	one.prev = current
	return one
}

type Day21DeterministicDice struct {
	NextValue int
	RollCount int
}

func NewDay21DeterministicDice() *Day21DeterministicDice {
	return &Day21DeterministicDice{
		NextValue: 1,
		RollCount: 0,
	}
}

func (d *Day21DeterministicDice) Roll() int {
	val := d.NextValue
	new := val + 1
	if new > 100 {
		new = 1
	}
	d.NextValue = new
	d.RollCount++
	return val
}

var day21Matcher = regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)

func parseDay21(lines []string) map[int]int {
	pos := map[int]int{}

	for _, l := range lines {
		matches := day21Matcher.FindStringSubmatch(l)
		if len(matches) != 3 {
			panic("did not match")
		}
		player := util.MustConvertDecimalStringToInt(matches[1])
		position := util.MustConvertDecimalStringToInt(matches[2])
		pos[player] = position
	}

	return pos
}
