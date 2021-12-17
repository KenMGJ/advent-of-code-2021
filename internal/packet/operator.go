package packet

import "fmt"

type Operator struct {
	idType     int
	lengthType int
	version    int
	subs       []Packet
}

func NewOperator(idType, lengthType, version int) *Operator {
	return &Operator{
		idType:     idType,
		lengthType: lengthType,
		version:    version,
		subs:       []Packet{},
	}
}

func (o *Operator) Add(p Packet) {
	o.subs = append(o.subs, p)
}

func (o *Operator) AddAll(pckts []Packet) {
	o.subs = append(o.subs, pckts...)
}

func (o *Operator) IdType() int {
	return o.idType
}

func (o *Operator) Subpackets() []Packet {
	return o.subs
}

func (o *Operator) Value() int {

	switch o.idType {
	case 0:
		// sum
		val := 0
		for _, s := range o.subs {
			val += s.Value()
		}
		return val
	case 1:
		// product
		val := 1
		for _, s := range o.subs {
			val *= s.Value()
		}
		return val
	case 2:
		// min
		min := -1
		for _, s := range o.subs {
			v := s.Value()
			if min == -1 || v < min {
				min = v
			}
		}
		return min
	case 3:
		// max
		max := -1
		for _, s := range o.subs {
			v := s.Value()
			if max == -1 || v > max {
				max = v
			}
		}
		return max
	case 5:
		// greater than
		if o.subs[0].Value() > o.subs[1].Value() {
			return 1
		} else {
			return 0
		}
	case 6:
		// less than
		if o.subs[0].Value() < o.subs[1].Value() {
			return 1
		} else {
			return 0
		}
	case 7:
		// equal to
		if o.subs[0].Value() == o.subs[1].Value() {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func (o *Operator) Version() int {
	return o.version
}

func (o *Operator) VersionSum() int {
	sum := o.version
	for _, o := range o.subs {
		sum += o.VersionSum()
	}

	return sum
}

func (o *Operator) String() string {
	return fmt.Sprintf("L{%d %d %d %v}", o.idType, o.lengthType, o.version, o.subs)
}
