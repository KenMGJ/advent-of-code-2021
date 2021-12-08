package runner

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/sets"
)

func (r *Runner) Day08Part1(lines []string) {
	input := parseDay08Input(lines)
	// fmt.Println(input)

	sum := 0
	for _, in := range input {
		icount := make(map[int]int)
		for _, i := range in.Input {
			icount[len(i)] += 1
		}

		// fmt.Println(icount)
		for _, o := range in.Output {
			if icount[len(o)] == 1 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func (r *Runner) Day08Part2(lines []string) {

	sum := 0

	input := parseDay08Input(lines)
	for _, in := range input {

		v2sMap := makeSignalToValueMap(in.Input)
		// fmt.Println(rem)
		// fmt.Println(v2sMap)

		signalToSegmentSetMap := make(map[string]*sets.StringSet)
		for _, r := range in.Input {
			letterSet := sets.NewStringSet()
			for _, s := range strings.Split(r, "") {
				letterSet.Add(s)
			}
			signalToSegmentSetMap[r] = letterSet
		}
		// fmt.Println(signalToSegmentSetMap)

		counts := make(map[string]int)
		for _, v := range signalToSegmentSetMap {
			for _, vv := range v.Vals() {
				counts[vv] += 1
			}
		}
		// fmt.Println(counts)

		segments := [7]string{}

		// segment 0
		n1 := v2sMap[1]
		s1 := signalToSegmentSetMap[n1]
		n2 := v2sMap[7]
		s2Vals := signalToSegmentSetMap[n2].Vals()

		// segment 0
		for _, v := range s2Vals {
			if !s1.Contains(v) {
				segments[0] = v
				break
			}
		}
		delete(counts, segments[0])

		// segment 1
		for k, v := range counts {
			if v == 6 {
				segments[1] = k
				break
			}
		}
		delete(counts, segments[1])

		// segment 2
		for k, v := range counts {
			if v == 8 {
				segments[2] = k
				break
			}
		}
		delete(counts, segments[2])

		// segment 4
		for k, v := range counts {
			if v == 4 {
				segments[4] = k
				break
			}
		}
		delete(counts, segments[4])

		// segment 5
		for k, v := range counts {
			if v == 9 {
				segments[5] = k
				break
			}
		}
		delete(counts, segments[5])

		// segment 3
		for k, v := range counts {
			if v == 7 && signalToSegmentSetMap[v2sMap[4]].Contains(k) {
				segments[3] = k
				break
			}
		}
		delete(counts, segments[3])

		for k := range counts {
			segments[6] = k
		}

		s := segments
		iToSig := [10]string{}
		iToSig[0] = normalizeDay08String(s[0] + s[1] + s[2] + s[4] + s[5] + s[6])
		iToSig[1] = normalizeDay08String(s[2] + s[5])
		iToSig[2] = normalizeDay08String(s[0] + s[2] + s[3] + s[4] + s[6])
		iToSig[3] = normalizeDay08String(s[0] + s[2] + s[3] + s[5] + s[6])
		iToSig[4] = normalizeDay08String(s[1] + s[2] + s[3] + s[5])
		iToSig[5] = normalizeDay08String(s[0] + s[1] + s[3] + s[5] + s[6])
		iToSig[6] = normalizeDay08String(s[0] + s[1] + s[3] + s[4] + s[5] + s[6])
		iToSig[7] = normalizeDay08String(s[0] + s[2] + s[5])
		iToSig[8] = normalizeDay08String(s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6])
		iToSig[9] = normalizeDay08String(s[0] + s[1] + s[2] + s[3] + s[5] + s[6])

		sigToI := make(map[string]int)
		for i, s := range iToSig {
			sigToI[s] = i
		}
		// fmt.Println(iToSig)

		num := ""
		for _, o := range in.Output {
			num = fmt.Sprintf("%s%d", num, sigToI[o])
		}
		numAsInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sum += numAsInt
		//fmt.Println(num)
	}

	fmt.Println(sum)
}

func makeSignalToValueMap(signals []string) map[int]string {
	valToSig := make(map[int]string)

	for i := 0; i < len(signals); i++ {
		sig := signals[i]
		switch len(sig) {
		case 2:
			valToSig[1] = sig
		case 3:
			valToSig[7] = sig
		case 4:
			valToSig[4] = sig
		case 7:
			valToSig[8] = sig
		}
	}

	return valToSig
}

func removeDay08SignalFromSlice(signals []string, idx int) []string {
	copy(signals[idx:], signals[idx+1:])
	signals[len(signals)-1] = ""
	return signals[:len(signals)-1]
}

type day08Data struct {
	Input  []string
	Output []string
}

func parseDay08Input(lines []string) []day08Data {
	input := []day08Data{}

	for _, l := range lines {
		parts := strings.Split(l, "|")

		input = append(input, day08Data{
			Input:  parseDay08PartToStringArray(parts[0]),
			Output: parseDay08PartToStringArray(parts[1]),
		})
	}

	return input
}

func parseDay08PartToStringArray(input string) []string {

	in := strings.Split(input, " ")
	new := []string{}
	for _, i := range in {
		s := strings.Trim(i, " ")
		if s != "" && s != " " {
			new = append(new, normalizeDay08String(s))
		}
	}

	return new
}

func normalizeDay08String(input string) string {
	in := strings.Split(input, "")
	sort.Strings(in)
	return strings.Join(in, "")
}
