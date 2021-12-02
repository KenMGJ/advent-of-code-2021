package main

import (
	"flag"
	"log"

	"github.com/KenMGJ/advent-of-code-2021/internal/runner"
)

func main() {

	dayFlag := flag.Int("day", 1, "day to run")
	testFlag := flag.Bool("test", false, "use test input")
	flag.Parse()

	runner := runner.New()
	err := runner.Run(*dayFlag, *testFlag)

	if err != nil {
		log.Fatal(err)
	}
}
