package runner

import "fmt"

type GameResult struct {
	P1Wins int
	P2Wins int
}

const WINNING_SCORE_DAY_21_PART_2 = 21

var gameCache map[string]GameResult

func Day21EndPos(pos, roll int) int {
	return ((pos - 1 + roll) % 10) + 1
}

// Reworked with a little help from:
// https://www.codingnagger.com/2021/12/22/the-advent-of-code-2021-day-21-late-night-finish/
func Day21Part2Run(p1pos, p1score, p2pos, p2score, turn, rollSum int) GameResult {
	isP1 := turn < 3
	lastP1Throw := turn == 2
	lastP2Throw := turn == 5

	cacheKey := fmt.Sprintf("%d,%d,%d,%d,%d,%d", p1pos, p1score, p2pos, p2score, turn, rollSum)

	var result GameResult
	if p1score >= WINNING_SCORE_DAY_21_PART_2 {
		result = GameResult{P1Wins: 1, P2Wins: 0}
	} else if p2score >= WINNING_SCORE_DAY_21_PART_2 {
		result = GameResult{P1Wins: 0, P2Wins: 1}
	} else {

		gr, ok := gameCache[cacheKey]
		if ok {
			return gr
		}

		nextTurn := (turn + 1) % 6

		p1Wins := 0
		p2Wins := 0

		for roll := 1; roll <= 3; roll++ {
			newp1score := p1score
			newp1pos := p1pos
			if isP1 {
				newp1pos = Day21EndPos(p1pos, roll)
				if lastP1Throw {
					newp1score += newp1pos
				}
			}

			newp2score := p2score
			newp2pos := p2pos
			if !isP1 {
				newp2pos = Day21EndPos(p2pos, roll)
				if lastP2Throw {
					newp2score += newp2pos
				}
			}

			nextRoll := 0
			if turn%3 != 0 {
				nextRoll = rollSum + roll
			}

			res := Day21Part2Run(newp1pos, newp1score, newp2pos, newp2score, nextTurn, nextRoll)

			p1Wins += res.P1Wins
			p2Wins += res.P2Wins
		}

		result = GameResult{P1Wins: p1Wins, P2Wins: p2Wins}

	}

	gameCache[cacheKey] = result
	return result
}
