package runner

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day24Part1(lines []string) {
	alu := NewALU()
	alu.Inp("w")
	fmt.Println(alu)
}

func (r *Runner) Day24Part2(lines []string) {
}

type ALU struct {
	mem map[string]int
}

func NewALU() *ALU {
	return &ALU{
		mem: map[string]int{},
	}
}

func (a *ALU) Inp(a0 string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s:", a0)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("error reading input")
	}
	text = strings.Trim(text, "\n")
	i := util.MustConvertDecimalStringToInt(text)
	a.mem[a0] = i
}

func (a *ALU) Add(a0, b0 string) {
	a.mem[a0] += a.mem[b0]
}

func (a *ALU) Mul(a0, b0 string) {
	a.mem[a0] *= a.mem[b0]
}

func (a *ALU) Div(a0, b0 string) {
	a.mem[a0] /= a.mem[b0]
}

func (a *ALU) Mod(a0, b0 string) {
	a.mem[a0] %= a.mem[b0]
}

func (a *ALU) Eql(a0, b0 string) {
	if a.mem[a0] == a.mem[b0] {
		a.mem[a0] = 1
	} else {
		a.mem[a0] = 0
	}
}
