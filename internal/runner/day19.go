package runner

import (
	"fmt"
	"math"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/point"
	"github.com/KenMGJ/advent-of-code-2021/internal/sets"
	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day19Part1(lines []string) {
	regions := parseDay19(lines)
	found := normalize(regions)

	fset := sets.NewStringSet()
	for _, f := range found {
		for _, b := range f.Beacons {
			fset.Add(b.String())
		}
	}
	fmt.Println(fset.Size())
}

func (r *Runner) Day19Part2(lines []string) {
	regions := parseDay19(lines)
	found := normalize(regions)

	max := float64(0)
	for i := 0; i < len(found)-1; i++ {
		for j := i + 1; j < len(found); j++ {

			absX := math.Abs(float64(found[i].Location.X() - found[j].Location.X()))
			absY := math.Abs(float64(found[i].Location.Y() - found[j].Location.Y()))
			absZ := math.Abs(float64(found[i].Location.Z() - found[j].Location.Z()))

			dist := absX + absY + absZ
			if dist > max {
				max = dist
			}
		}
	}
	fmt.Println(max)
}

func normalize(regions []*ScannerRegion) []*ScannerRegion {

	found := []*ScannerRegion{}
	r0 := regions[0]
	r0.Location = *point.NewPoint3D(0, 0, 0)
	found = append(found, r0)
	regions = regions[1:]

	for len(regions) > 0 {
		for _, f := range found {

			fset := sets.NewFloat64Set()
			for d := range f.DistPointMap() {
				fset.Add(d)
			}

			matchFound := false
			for i, r := range regions {
				rset := sets.NewFloat64Set()
				for d := range r.DistPointMap() {
					rset.Add(d)
				}

				for j := 0; j <= 24; j++ {

					intrset := fset.Intersect(rset)
					if intrset.Size() >= 66 {

						anyDistPointKey := intrset.Vals()[1]
						f0 := f.DistPointMap()[anyDistPointKey]
						r0 := r.DistPointMap()[anyDistPointKey]

						dAX := f0.A.X() - r0.A.X()
						dAY := f0.A.Y() - r0.A.Y()
						dAZ := f0.A.Z() - r0.A.Z()
						dBX := f0.B.X() - r0.B.X()
						dBY := f0.B.Y() - r0.B.Y()
						dBZ := f0.B.Z() - r0.B.Z()

						if dAX == dBX && dAY == dBY && dAZ == dBZ {

							for _, b := range r.Beacons {
								b.Set(b.X()+dAX, b.Y()+dAY, b.Z()+dAZ)
							}

							matchFound = true
							r.Location = *point.NewPoint3D(dAX, dAY, dAZ)
							found = append(found, r)
							regions = append(regions[:i], regions[i+1:]...)

							break
						}

						dAX = f0.A.X() - r0.B.X()
						dAY = f0.A.Y() - r0.B.Y()
						dAZ = f0.A.Z() - r0.B.Z()
						dBX = f0.B.X() - r0.A.X()
						dBY = f0.B.Y() - r0.A.Y()
						dBZ = f0.B.Z() - r0.A.Z()

						if dAX == dBX && dAY == dBY && dAZ == dBZ {

							for _, b := range r.Beacons {
								b.Set(b.X()+dAX, b.Y()+dAY, b.Z()+dAZ)
							}

							matchFound = true
							r.Location = *point.NewPoint3D(dAX, dAY, dAZ)
							found = append(found, r)
							regions = append(regions[:i], regions[i+1:]...)

							break
						}
					}

					r.Rotate()
				}

				if matchFound {
					break
				}
			}
		}
	}

	return found
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

func (s *ScannerRegion) DistPointMap() map[float64]*Point3DPair {
	mapP2P := map[float64]*Point3DPair{}

	for i := 0; i < len(s.Beacons)-1; i++ {
		for j := i + 1; j < len(s.Beacons); j++ {

			dX := s.Beacons[i].X() - s.Beacons[j].X()
			dX *= dX

			dY := s.Beacons[i].Y() - s.Beacons[j].Y()
			dY *= dY

			dZ := s.Beacons[i].Z() - s.Beacons[j].Z()
			dZ *= dZ

			sum := dX + dY + dZ
			dist := math.Sqrt(float64(sum))

			mapP2P[dist] = &Point3DPair{
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
