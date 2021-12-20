package point

import "math"

type Point3D struct {
	X int
	Y int
	Z int
}

var origin = Point3D{X: 0, Y: 0, Z: 0}

// See the top comment:
// https://stackoverflow.com/questions/16452383/how-to-get-all-24-rotations-of-a-3-dimensional-array
func (c Point3D) Rotations() []Point3D {
	rots := []Point3D{}

	for cycle := 0; cycle < 2; cycle++ {
		for step := 0; step < 3; step++ {
			c = c.roll()
			rots = append(rots, c)
			for i := 0; i < 3; i++ {
				c = c.turn()
				rots = append(rots, c)
			}
		}
		c = c.roll().turn().roll()
	}

	return rots
}

func (c Point3D) DistanceFromOrigin() float64 {
	return c.DistanceFrom(origin)
}

func (c Point3D) DistanceFrom(o Point3D) float64 {
	dX2 := math.Pow(float64(c.X-o.X), 2)
	dY2 := math.Pow(float64(c.Y-o.Y), 2)
	dZ2 := math.Pow(float64(c.Z-o.Z), 2)

	return math.Sqrt(dX2 + dY2 + dZ2)
}

func (p Point3D) Equals(o Point3D) bool {
	return p.X == o.X && p.Y == o.Y && p.Z == o.Z
}

func Rotations(points []Point3D) [][]Point3D {
	rots := [][]Point3D{}

	for i := 0; i < len(points); i++ {
		row := []Point3D{}
		for j := 0; j < 24; j++ {
			row = append(row, Point3D{})
		}
		rots = append(rots, row)
	}

	for i, p := range points {
		r := p.Rotations()
		for j, p2 := range r {
			rots[i][j] = p2
		}
	}

	return rots
}

func (p Point3D) roll() Point3D {
	return Point3D{
		X: p.X,
		Y: p.Z,
		Z: -1 * p.Y,
	}
}

func (p Point3D) turn() Point3D {
	return Point3D{
		X: -1 * p.Y,
		Y: p.X,
		Z: p.Z,
	}
}
