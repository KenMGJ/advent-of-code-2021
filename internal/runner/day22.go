package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day22Part1(lines []string) {
	steps := parseDay22(lines)

	core := NewReactorCore()

	i := 1
	for _, s := range steps {

		xmin := util.MaxInt(-50, s.XMin)
		ymin := util.MaxInt(-50, s.YMin)
		zmin := util.MaxInt(-50, s.ZMin)

		xmax := util.MinInt(50, s.XMax)
		ymax := util.MinInt(50, s.YMax)
		zmax := util.MinInt(50, s.ZMax)

		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				for z := zmin; z <= zmax; z++ {
					core.Set(x, y, z, s.On)
				}
			}
		}
		i++
	}

	count := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if core.Get(x, y, z) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func (r *Runner) Day22Part2(lines []string) {
	steps := parseDay22(lines)

	for _, s := range steps {
		fmt.Println(s.Volume())
	}
}

type ReactorCore struct {
	core map[int]map[int]map[int]bool
}

func NewReactorCore() *ReactorCore {
	return &ReactorCore{
		core: map[int]map[int]map[int]bool{},
	}
}

func (r *ReactorCore) Get(x, y, z int) bool {

	xCore, ok := r.core[x]
	if !ok {
		xCore = map[int]map[int]bool{}
	}

	yCore, ok := xCore[y]
	if !ok {
		yCore = map[int]bool{}
	}

	val, ok := yCore[z]
	if !ok {
		val = false
	}

	yCore[z] = val
	xCore[y] = yCore
	r.core[x] = xCore

	return val
}

func (r *ReactorCore) Set(x, y, z int, on bool) {

	xCore, ok := r.core[x]
	if !ok {
		xCore = map[int]map[int]bool{}
	}

	yCore, ok := xCore[y]
	if !ok {
		yCore = map[int]bool{}
	}

	yCore[z] = on
	xCore[y] = yCore
	r.core[x] = xCore
}

var matcherDay22 = regexp.MustCompile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)

type RebootStep struct {
	On   bool
	XMin int
	XMax int
	YMin int
	YMax int
	ZMin int
	ZMax int
}

func (r *RebootStep) Volume() int {
	return (r.XMax - r.XMin) * (r.YMax - r.YMin) * (r.ZMax - r.ZMin)
}

func parseDay22(lines []string) []RebootStep {
	steps := []RebootStep{}

	for _, l := range lines {
		matches := matcherDay22.FindStringSubmatch(l)

		if len(matches) != 8 {
			panic("invalid format")
		}

		on := true
		if matches[1] == "off" {
			on = false
		}

		steps = append(steps, RebootStep{
			On:   on,
			XMin: util.MustConvertDecimalStringToInt(matches[2]),
			XMax: util.MustConvertDecimalStringToInt(matches[3]),
			YMin: util.MustConvertDecimalStringToInt(matches[4]),
			YMax: util.MustConvertDecimalStringToInt(matches[5]),
			ZMin: util.MustConvertDecimalStringToInt(matches[6]),
			ZMax: util.MustConvertDecimalStringToInt(matches[7]),
		})
	}

	return steps
}
