package runner

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/packet"
)

func (r *Runner) Day16Part1(lines []string) {
	for _, line := range lines {
		if strings.HasPrefix(line, "--") {
			continue
		}

		data := mustParseDay16Line(line)
		p, _ := processDay16Packets(data)
		fmt.Println(p.VersionSum())
	}
}

func (r *Runner) Day16Part2(lines []string) {
	for _, line := range lines {
		if strings.HasPrefix(line, "--") {
			continue
		}

		data := mustParseDay16Line(line)
		p, _ := processDay16Packets(data)
		// fmt.Println(p)
		fmt.Println(p.Value())
	}
}

func processDay16Packets(bits string) (packet.Packet, string) {

	version, b, ok := extractVersion(bits)
	if !ok {
		return nil, ""
	}
	bits = b

	idType, b, ok := extractIdType(bits)
	if !ok {
		return nil, ""
	}
	bits = b

	if idType == 4 {
		l, b := processDay16Literal(bits, idType, version)
		return l, b

	} else {

		lengthType, b := extractLengthType(bits)
		bits = b

		if lengthType == 0 {

			length, b2, ok := extractLengthType0(bits)
			if !ok {
				return nil, ""
			}
			bits = b2[:length]
			leftover := b2[length:]

			p := packet.NewOperator(idType, lengthType, version)
			for len(bits) > 0 {
				s, b := processDay16Packets(bits)
				bits = b
				p.Add(s)
			}
			return p, leftover

		} else {

			length, b, ok := extractLengthType1(bits)
			if !ok {
				return nil, ""
			}
			bits = b

			p := packet.NewOperator(idType, lengthType, version)

			for i := 0; i < length; i++ {
				s, b := processDay16Packets(bits)
				bits = b
				p.Add(s)
			}

			return p, bits
		}

	}
}

func processDay16Literal(bits string, idType, version int) (*packet.LiteralValue, string) {

	value := ""

	for {

		purpose := bits[0]
		bits = bits[1:]

		next := bits[:4]
		bits = bits[4:]

		value += next

		if mustConvertBinaryStringToInt(string(purpose)) == 0 {
			break
		}
	}

	val := mustConvertBinaryStringToInt(value)
	return packet.NewLiteralValue(idType, val, version), bits
}

func extractVersion(bits string) (int, string, bool) {
	if len(bits) < 3 {
		return 0, "", false
	}
	versionStr := bits[:3]
	return mustConvertBinaryStringToInt(versionStr), bits[3:], true
}

func extractIdType(bits string) (int, string, bool) {
	if len(bits) < 3 {
		return 0, "", false
	}
	idTypeStr := bits[:3]
	return mustConvertBinaryStringToInt(idTypeStr), bits[3:], true
}

func extractLengthType(bits string) (int, string) {
	lengthTypeStr := bits[:1]
	return mustConvertBinaryStringToInt(string(lengthTypeStr)), bits[1:]
}

func extractLengthType0(bits string) (int, string, bool) {
	if len(bits) < 15 {
		return 0, "", false
	}
	return mustConvertBinaryStringToInt(bits[:15]), bits[15:], true
}

func extractLengthType1(bits string) (int, string, bool) {
	if len(bits) < 11 {
		return 0, "", false
	}
	lengthBits := bits[:11]
	return mustConvertBinaryStringToInt(lengthBits), bits[11:], true
}

func mustConvertBinaryStringToInt(bin string) int {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func mustParseDay16Line(line string) string {

	bin := ""

	for _, c := range strings.Split(line, "") {
		h, err := strconv.ParseInt(c, 16, 64)
		if err != nil {
			panic(err)
		}
		bin += fmt.Sprintf("%04b", h)
	}

	return bin
}
