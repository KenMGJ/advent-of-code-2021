package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/point"
	"github.com/KenMGJ/advent-of-code-2021/internal/sets"
	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day19Part1(lines []string) {
	regions := parseDay19(lines)

	r0 := regions[0]
	r0set := sets.NewStringSet()
	for k := range r0.DistPointMap() {
		r0set.Add(string(k))
	}

	i := 1
	for j := 0; j <= 24; j++ {
		r1 := regions[i]
		r1set := sets.NewStringSet()
		for k := range r1.DistPointMap() {
			r1set.Add(string(k))
		}

		intr0r1 := r0set.Intersect(r1set)
		if intr0r1.Size() >= 66 {
			fmt.Println(i)
			fmt.Println(intr0r1)

			anyDistPointKey := DistPointKey(intr0r1.Vals()[0])
			p0 := r0.DistPointMap()[anyDistPointKey]
			p1 := r1.DistPointMap()[anyDistPointKey]

			fmt.Println(p0)
			fmt.Println(p1)

			break
		}

		r1.Rotate()
	}

}

func (r *Runner) Day19Part2(lines []string) {
}

type Point3DPair struct {
	A *point.Point3D
	B *point.Point3D
}

func (p *Point3DPair) String() string {
	return fmt.Sprintf("(%d,%d,%d) -> (%d,%d,%d)", p.A.X(), p.A.Y(), p.A.Z(), p.B.X(), p.B.Y(), p.B.Z())
}

type ScannerRegion struct {
	ID       int
	Location point.Point3D
	Beacons  []*point.Point3D
}

type DistPointKey string

var distPointKeyMatcher = regexp.MustCompile(`\((\d+),(\d+),(\d+)\)`)

func (d DistPointKey) Parts() (int, int, int) {

	matches := distPointKeyMatcher.FindStringSubmatch(string(d))
	if len(matches) != 4 {
		panic("at the disco")
	}

	x := util.MustConvertDecimalStringToInt(matches[1])
	y := util.MustConvertDecimalStringToInt(matches[2])
	z := util.MustConvertDecimalStringToInt(matches[3])

	return x, y, z
}

func (s *ScannerRegion) DistPointMap() map[DistPointKey]*Point3DPair {
	mapP2P := map[DistPointKey]*Point3DPair{}

	for i := 0; i < len(s.Beacons)-1; i++ {
		for j := i + 1; j < len(s.Beacons); j++ {

			absX := s.Beacons[i].X() - s.Beacons[j].X()
			if absX < 0 {
				absX *= -1
			}

			absY := s.Beacons[i].Y() - s.Beacons[j].Y()
			if absY < 0 {
				absY *= -1
			}

			absZ := s.Beacons[i].Z() - s.Beacons[j].Z()
			if absZ < 0 {
				absZ *= -1
			}

			mapP2P[DistPointKey(fmt.Sprintf("(%d,%d,%d)", absX, absY, absZ))] = &Point3DPair{
				A: s.Beacons[i],
				B: s.Beacons[j],
			}

		}
	}

	return mapP2P
}

func (s *ScannerRegion) Rotate() {
	for _, b := range s.Beacons {
		b.Rotate()
	}
}

func (s *ScannerRegion) String() string {
	return fmt.Sprintf("{%d %v}", s.ID, s.Beacons)
}

var scannerMatcher = regexp.MustCompile(`^--- scanner (\d+) ---$`)
var coordinate3DMatcher = regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+)$`)

func parseDay19(lines []string) []*ScannerRegion {
	scanners := []*ScannerRegion{}

	var scanner *ScannerRegion

	for _, l := range lines {

		scannerParts := scannerMatcher.FindStringSubmatch(l)
		c3Dparts := coordinate3DMatcher.FindStringSubmatch(l)

		if len(scannerParts) == 2 {

			if scanner != nil {
				scanners = append(scanners, scanner)
			}

			scanner = &ScannerRegion{
				ID:      util.MustConvertDecimalStringToInt(scannerParts[1]),
				Beacons: []*point.Point3D{},
			}

		} else if len(c3Dparts) == 4 {

			scanner.Beacons = append(scanner.Beacons,
				point.NewPoint3D(
					util.MustConvertDecimalStringToInt(c3Dparts[1]),
					util.MustConvertDecimalStringToInt(c3Dparts[2]),
					util.MustConvertDecimalStringToInt(c3Dparts[3]),
				))
		}
	}

	scanners = append(scanners, scanner)
	return scanners
}
