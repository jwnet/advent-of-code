package main

import (
	"fmt"
	"strconv"
	"strings"
)

type packetData struct {
	hexStr string
	bitBuffer     uint64
	nBufferedBits uint64
}

func newPacketData(hex string) *packetData {
	return &packetData{hexStr: hex}
}

func (p *packetData) nextByte() byte {
	if len(p.hexStr) < 2 {
		panic("packetData.nextByte: trying to read into something that's not there...")
	}
	b, err := strconv.ParseUint(p.hexStr[:2], 16, 64)
	if err != nil {
		panic("packetData.nextByte: your damned input is invalid")
	}
	p.hexStr = p.hexStr[2:]
	return byte(b)
}

// NextNBits returns up to 64 bits at a time. Panics if you request
// more than 64 bits. Behave. We also panic if you run out of packet
// data to read. Yikes!
func (p *packetData) NextNBits(n uint64) uint64 {
	for p.nBufferedBits < n {
		if p.nBufferedBits+8 > 64 {
			panic("packetData.NextNBits: you requested more than 64 bits")
		}
		p.bitBuffer = (p.bitBuffer << 8) + uint64(p.nextByte())
		p.nBufferedBits += 8
	}
	b := p.bitBuffer >> (p.nBufferedBits - n)
	adjust := 64 + n - p.nBufferedBits
	p.bitBuffer = (p.bitBuffer << adjust) >> adjust
	p.nBufferedBits -= n
	return b
}

//---------------------------------------------------------

const (
	id_sum  = 0
	id_prod = 1 // product of sub-packets or single packet
	id_min  = 2 // min of sub-packets
	id_max  = 3 // max of sub-packets
	id_lit  = 4 // literal value
	id_gt   = 5 // 1 if first packet > second packet
	id_lt   = 6 // 1 if first packet < second packet
	id_eq   = 7 // 1 if equal
	len_15  = 0 // 15-bit number specifying number of bits in sub-packets
	len_11  = 1 // 11-bit number specifying number of sub-packets
)

func sum(s []uint64) uint64 {
	var sum uint64 = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func prod(s []uint64) uint64 {
	var prod uint64 = 1
	for _, v := range s {
		prod *= v
	}
	return prod
}

func min(s []uint64) uint64 {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func max(s []uint64) uint64 {
	var max uint64 = 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func gt(s []uint64) uint64 {
	if s[0] > s[1] {
		return 1
	}
	return 0
}

func lt(s []uint64) uint64 {
	if s[0] < s[1] {
		return 1
	}
	return 0
}

func eq(s []uint64) uint64 {
	if s[0] == s[1] {
		return 1
	}
	return 0
}

type packet struct {
	nBits       uint64
	version, id uint64
	literal     uint64
	lengthType  uint64
	sub         []*packet
}

func (p *packet) String(prefix string) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s packet:\n", prefix))
	b.WriteString(fmt.Sprintf("%s version: %d\n", prefix, p.version))
	b.WriteString(fmt.Sprintf("%s      id: %d\n", prefix, p.id))
	b.WriteString(fmt.Sprintf("%slengthID: %d\n", prefix, p.lengthType))
	if len(p.sub) > 0 {
		b.WriteString(fmt.Sprintf("%s%d sub-packets: \n", prefix, len(p.sub)))
		for _, sp := range p.sub {
			b.WriteString(sp.String(prefix + "\t"))
		}
	}
	return b.String()
}

func (p *packet) value() (val uint64) {
	if p.id == id_lit {
		return p.literal
	}
	subVals := make([]uint64, len(p.sub))
	for i, sp := range p.sub {
		subVals[i] = sp.value()
	}
	switch p.id {
	case id_sum:
		val = sum(subVals)
	case id_prod:
		val = prod(subVals)
	case id_min:
		val = min(subVals)
	case id_max:
		val = max(subVals)
	case id_gt:
		val = gt(subVals)
	case id_lt:
		val = lt(subVals)
	case id_eq:
		val = eq(subVals)
	}
	return
}

func (p *packet) parse(pd *packetData) {
	p.version = p.takeNBits(pd, 3)
	p.id = p.takeNBits(pd, 3)
	switch p.id {
	case id_lit:
		p.parseLiteral(pd)
	default:
		p.parseOperator(pd)
	}
}

func (p *packet) parseLiteral(pd *packetData) {
	var literal uint64 = 0
	for end := false; !end; {
		end = p.takeNBits(pd, 1) == 0
		literal = (literal << (4)) + p.takeNBits(pd, 4)
	}
	p.literal = literal
}

func (p *packet) parseOperator(pd *packetData) {
	p.lengthType = p.takeNBits(pd, 1)
	switch p.lengthType {
	case len_15:
		p.parseL15(pd)
	case len_11:
		p.parseL11(pd)
	}
}

func (p *packet) parseL11(pd *packetData) {
	npackets := p.takeNBits(pd, 11)
	var used uint64
	for i := uint64(0); i < npackets; i++ {
		np := new(packet)
		np.parse(pd)
		p.sub = append(p.sub, np)
		used += np.nBits
	}
	p.nBits += used
}

func (p *packet) parseL15(pd *packetData) {
	nbits := p.takeNBits(pd, 15)
	var used uint64
	for used < nbits {
		np := new(packet)
		np.parse(pd)
		p.sub = append(p.sub, np)
		used += np.nBits
	}
	p.nBits += used
}

func (p *packet) versionSum() uint64 {
	v := p.version
	for _, sp := range p.sub {
		v += sp.versionSum()
	}
	return v
}

func (p *packet) takeNBits(pd *packetData, n uint64) uint64 {
	p.nBits += n
	return pd.NextNBits(n)
}

func part1(packetBlob string) uint64 {
	pd := newPacketData(packetBlob)
	p := new(packet)
	p.parse(pd)
	return p.versionSum()
}

func part2(packetBlob string) uint64 {
	pd := newPacketData(packetBlob)
	p := new(packet)
	p.parse(pd)
	return p.value()
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample1: %d\n", part1("D2FE28"))                         // 4
	fmt.Printf("\tsample2: %d\n", part1("38006F45291200"))                 // 9
	fmt.Printf("\tsample3: %d\n", part1("EE00D40C823060"))                 // 14
	fmt.Printf("\tsample4: %d\n", part1("8A004A801A8002F478"))             // 16
	fmt.Printf("\tsample5: %d\n", part1("620080001611562C8802118E34"))     // 12
	fmt.Printf("\tsample6: %d\n", part1("C0015000016115A2E0802F182340"))   // 23
	fmt.Printf("\tsample7: %d\n", part1("A0016C880162017C3686B18A3D4780")) // 31
	fmt.Printf("\t  input: %d\n", part1(input))                            // 951
	fmt.Println("Part 2")
	fmt.Printf("\t sample8: %d\n", part2("C200B40A82"))                 // 3
	fmt.Printf("\t sample9: %d\n", part2("04005AC33890"))               // 54
	fmt.Printf("\tsample10: %d\n", part2("880086C3E88112"))             // 7
	fmt.Printf("\tsample11: %d\n", part2("CE00C43D881120"))             // 9
	fmt.Printf("\tsample12: %d\n", part2("D8005AC2A8F0"))               // 1
	fmt.Printf("\tsample13: %d\n", part2("F600BC2D8F"))                 // 0
	fmt.Printf("\tsample14: %d\n", part2("9C005AC2F8F0"))               // 0
	fmt.Printf("\tsample15: %d\n", part2("9C0141080250320F1802104A08")) // 1
	fmt.Printf("\t   input: %d\n", part2(input))                        // 902198718880
}
