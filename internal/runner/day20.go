package runner

import (
	"fmt"
	"strings"
)

func (r *Runner) Day20Part1(lines []string) {
	img := parseDay20(lines)
	execDay20(img, 2)
}

func (r *Runner) Day20Part2(lines []string) {
	img := parseDay20(lines)
	execDay20(img, 50)
}

func execDay20(img *Image, iterations int) {
	for i := 0; i < iterations; i++ {
		img.Enhance()
	}
	fmt.Println(img.LitCount())
}

func parseDay20(lines []string) *Image {
	algStr := lines[0]
	alg := []bool{}
	for _, c := range strings.Split(algStr, "") {
		if c == "#" {
			alg = append(alg, true)
		} else {
			alg = append(alg, false)
		}
	}

	lines = lines[2:]

	img := &Image{
		Alg:     alg,
		Default: false,
		Img:     map[int]map[int]bool{},
	}
	for i := 0; i < len(lines); i++ {
		for j, c := range strings.Split(lines[i], "") {
			if c == "#" {
				img.Put(i, j, true)
			} else {
				img.Put(i, j, false)
			}
		}
	}

	return img
}

type Image struct {
	Alg     []bool
	Default bool
	Img     map[int]map[int]bool
}

func (img *Image) Enhance() {

	minA, maxA, minB, maxB := img.getIndexes()

	newImg := &Image{
		Img: map[int]map[int]bool{},
	}

	for a := minA - 1; a <= maxA+1; a++ {
		for b := minB - 1; b <= maxB+1; b++ {
			bits := []bool{}

			for a1 := a - 1; a1 <= a+1; a1++ {
				for b1 := b - 1; b1 <= b+1; b1++ {
					bits = append(bits, img.Get(a1, b1))
				}
			}

			val := 0
			for _, bit := range bits {
				bt := 0
				if bit {
					bt = 1
				}
				val = val << 1
				val = val | bt
			}

			newImg.Put(a, b, img.Alg[val])
		}
	}

	img.Img = newImg.Img
	if img.Alg[0] {
		img.Default = !img.Default
	}
}

func (i *Image) Get(a, b int) bool {

	_, ok := i.Img[a]
	if !ok {
		i.Img[a] = map[int]bool{}
	}

	_, ok = i.Img[a][b]
	if !ok {
		i.Img[a][b] = i.Default
	}

	return i.Img[a][b]
}

func (i *Image) LitCount() int {

	lit := 0

	for _, l := range i.Img {
		for _, p := range l {
			if p {
				lit++
			}
		}
	}

	return lit
}

func (i *Image) Print() {
	minA, maxA, minB, maxB := i.getIndexes()

	for a := minA; a <= maxA; a++ {
		for b := minB; b <= maxB; b++ {
			bit := i.Get(a, b)
			if bit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (i *Image) Put(a, b int, value bool) {

	_, ok := i.Img[a]
	if !ok {
		i.Img[a] = map[int]bool{}
	}

	i.Img[a][b] = value

}

func (i *Image) getIndexes() (int, int, int, int) {
	minI, maxI, minJ, maxJ := 0, 0, 0, 0
	for k, v := range i.Img {
		if k < minI {
			minI = k
		} else if k > maxI {
			maxI = k
		}

		for k1 := range v {
			if k1 < minJ {
				minJ = k1
			} else if k1 > maxJ {
				maxJ = k1
			}
		}
	}

	return minI, maxI, minJ, maxJ
}
