package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/point"
	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day19Part1(lines []string) {
	beacons := parseDay19(lines)
	fmt.Println(beacons)

	// oBeacon := beacons[0]
	// oBeaconDistances := makeDistanceToPairMap(oBeacon)

	/*
		for i := 1; i < len(beacons); i++ {
			cmp := beacons[i]
			cmpAll := point.Rotations(cmp)
			for _, ca := range cmpAll {
				fmt.Println(ca)
				cmpDistances := makeDistanceToPairMap(ca)
				for k, v := range oBeaconDistances {

					cv, ok := cmpDistances[k]
					if ok {
						vxDiff := v[0][0].X - v[0][1].X
						fmt.Println(vxDiff)
						vyDiff := v[0][0].Y - v[0][1].Y
						fmt.Println(vyDiff)
						vzDiff := v[0][0].Z - v[0][1].Z
						fmt.Println(vzDiff)
						cvxDiff := cv[0][0].X - cv[0][1].X
						fmt.Println(cvxDiff)
						cvyDiff := cv[0][0].Y - cv[0][1].Y
						fmt.Println(cvyDiff)
						cvzDiff := cv[0][0].Z - cv[0][1].Z
						fmt.Println(cvzDiff)

						if vxDiff == cvxDiff && vyDiff == cvyDiff && vzDiff == cvzDiff {
							fmt.Println(k)
							fmt.Println(v)
							fmt.Println(cv)
							fmt.Println()
						}
					}
				}
			}
		}
	*/
}

func makeDistanceToPairMap(points []point.Point3D) map[float64][][]point.Point3D {
	distances := map[float64][][]point.Point3D{}

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pI := points[i]
			pJ := points[j]
			key := pI.DistanceFrom(pJ)
			v, ok := distances[key]
			if !ok {
				v = [][]point.Point3D{}
			}
			pts := []point.Point3D{}
			if pI.X >= pJ.X {
				pts = append(pts, pI, pJ)
			} else {
				pts = append(pts, pJ, pI)
			}
			v = append(v, pts)
			distances[key] = v
		}
	}

	return distances
}

func (r *Runner) Day19Part2(lines []string) {
	/*
		beacons := parseDay19(lines)
		fmt.Println(beacons[0][0])
		seq := beacons[0][0].Rotations()
		fmt.Println(seq)

		rots := point.Rotations(beacons[0])
		fmt.Println(rots)
	*/
}

type ScannerRegion struct {
	ID       int
	Location point.Point3D
	Beacons  []point.Point3D
}

func (s *ScannerRegion) String() string {
	return fmt.Sprintf("SR{ID: %d Count: %d}", s.ID, len(s.Beacons))
}

var scannerMatcher = regexp.MustCompile(`^--- scanner (\d+) ---$`)
var coordinate2DMatcher = regexp.MustCompile(`^(-?\d+),(-?\d+)$`)
var coordinate3DMatcher = regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+)$`)

func parseDay19(lines []string) []*ScannerRegion {
	scanners := []*ScannerRegion{}

	var scanner *ScannerRegion

	for _, l := range lines {

		scannerParts := scannerMatcher.FindStringSubmatch(l)
		c2DParts := coordinate2DMatcher.FindStringSubmatch(l)
		c3Dparts := coordinate3DMatcher.FindStringSubmatch(l)

		if len(scannerParts) == 2 {

			scanner = &ScannerRegion{
				ID:      util.MustConvertDecimalStringToInt(scannerParts[1]),
				Beacons: []point.Point3D{},
			}

		} else if len(c2DParts) == 3 {

			scanner.Beacons = append(scanner.Beacons, point.Point3D{
				X: util.MustConvertDecimalStringToInt(c2DParts[1]),
				Y: util.MustConvertDecimalStringToInt(c2DParts[2]),
			})

		} else if len(c3Dparts) == 4 {

			scanner.Beacons = append(scanner.Beacons, point.Point3D{
				X: util.MustConvertDecimalStringToInt(c3Dparts[1]),
				Y: util.MustConvertDecimalStringToInt(c3Dparts[2]),
				Z: util.MustConvertDecimalStringToInt(c3Dparts[3]),
			})

		} else {
			scanners = append(scanners, scanner)
		}
	}

	scanners = append(scanners, scanner)

	return scanners
}
