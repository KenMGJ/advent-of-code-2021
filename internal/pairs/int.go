package pairs

type IntPair struct {
	A int
	B int
}

func AdjacentPairs(row, column, height, width int) []IntPair {
	pairs := []IntPair{}

	if row > 0 {
		if column > 0 {
			pairs = append(pairs, IntPair{A: row - 1, B: column - 1})
		}
		pairs = append(pairs, IntPair{A: row - 1, B: column})
		if column < width-1 {
			pairs = append(pairs, IntPair{A: row - 1, B: column + 1})
		}
	}

	if column > 0 {
		pairs = append(pairs, IntPair{A: row, B: column - 1})
	}

	if column < width-1 {
		pairs = append(pairs, IntPair{A: row, B: column + 1})
	}

	if row < height-1 {
		if column > 0 {
			pairs = append(pairs, IntPair{A: row + 1, B: column - 1})
		}
		pairs = append(pairs, IntPair{A: row + 1, B: column})
		if column < width-1 {
			pairs = append(pairs, IntPair{A: row + 1, B: column + 1})
		}
	}

	return pairs
}
