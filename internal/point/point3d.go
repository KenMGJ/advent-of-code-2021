package point

import (
	"fmt"
	"math"
)

type Point3D struct {
	rotation int
	refX     int
	refY     int
	refZ     int
	x        int
	y        int
	z        int
}

func NewPoint3D(x, y, z int) *Point3D {
	return &Point3D{
		rotation: 0,
		refX:     x,
		refY:     y,
		refZ:     z,
		x:        x,
		y:        y,
		z:        z,
	}
}

func (p *Point3D) Distance(other *Point3D) float64 {
	return math.Sqrt(math.Pow(float64(p.x-other.x), 2) + math.Pow(float64(p.y-other.y), 2) + math.Pow(float64(p.z-other.z), 2))
}

func (p *Point3D) X() int { return p.x }
func (p *Point3D) Y() int { return p.y }
func (p *Point3D) Z() int { return p.z }

func (p *Point3D) Rotate() {

	// https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
	switch p.rotation {
	case 0:
		p.x, p.y, p.z = -1*p.refZ, p.refY, p.refX
	case 1:
		p.x, p.y, p.z = -1*p.refX, p.refY, -1*p.refZ
	case 2:
		p.x, p.y, p.z = p.refZ, p.refY, -1*p.refX
	case 3:
		p.x, p.y, p.z = p.refY, -1*p.refX, p.refZ
	case 4:
		p.x, p.y, p.z = p.refY, p.refZ, p.refX
	case 5:
		p.x, p.y, p.z = p.refY, p.refX, -1*p.refZ
	case 6:
		p.x, p.y, p.z = p.refY, -1*p.refZ, -1*p.refX
	case 7:
		p.x, p.y, p.z = -1*p.refY, p.refX, p.refZ
	case 8:
		p.x, p.y, p.z = -1*p.refY, -1*p.refZ, p.refX
	case 9:
		p.x, p.y, p.z = -1*p.refY, -1*p.refX, -1*p.refZ
	case 10:
		p.x, p.y, p.z = -1*p.refY, p.refZ, -1*p.refX
	case 11:
		p.x, p.y, p.z = p.refX, p.refZ, -1*p.refY
	case 12:
		p.x, p.y, p.z = -1*p.refZ, p.refX, -1*p.refY
	case 13:
		p.x, p.y, p.z = -1*p.refX, -1*p.refZ, -1*p.refY
	case 14:
		p.x, p.y, p.z = p.refZ, -1*p.refX, -1*p.refY
	case 15:
		p.x, p.y, p.z = p.refX, -1*p.refY, -1*p.refZ
	case 16:
		p.x, p.y, p.z = -1*p.refZ, -1*p.refY, -1*p.refX
	case 17:
		p.x, p.y, p.z = -1*p.refX, -1*p.refY, p.refZ
	case 18:
		p.x, p.y, p.z = p.refZ, -1*p.refY, p.refX
	case 19:
		p.x, p.y, p.z = p.refX, -1*p.refZ, p.refY
	case 20:
		p.x, p.y, p.z = -1*p.refZ, -1*p.refX, p.refY
	case 21:
		p.x, p.y, p.z = -1*p.refX, p.refZ, p.refY
	case 22:
		p.x, p.y, p.z = p.refZ, p.refX, p.refY
	case 23:
		p.x, p.y, p.z = p.refX, p.refY, p.refZ
	}

	if p.rotation == 23 {
		p.rotation = 0
	} else {
		p.rotation++
	}
}

func (p *Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}
