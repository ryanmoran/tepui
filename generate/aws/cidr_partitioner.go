package aws

import (
	"encoding/binary"
	"math"
	"net"
)

type CIDRPartitioner struct {
	CIDR  string
	Parts int
}

func NewCIDRPartitioner(cidr string, parts int) CIDRPartitioner {
	return CIDRPartitioner{
		CIDR:  cidr,
		Parts: parts,
	}
}

func (p CIDRPartitioner) Partition(index int) string {
	_, ipNet, err := net.ParseCIDR(p.CIDR)
	if err != nil {
		panic(err)
	}

	ones, bits := ipNet.Mask.Size()
	partition := int(math.Pow(2, float64(bits-ones))) / nextPowerOf2(p.Parts)
	ones = bits - int(math.Log2(float64(partition)))

	ip4 := ipNet.IP.To4()
	n := binary.BigEndian.Uint32(ip4)
	n += uint32(index) << uint(bits-ones)

	subnetIP := make(net.IP, len(ip4))
	binary.BigEndian.PutUint32(subnetIP, n)

	partitionNet := &net.IPNet{
		IP:   subnetIP,
		Mask: net.CIDRMask(ones, bits),
	}

	return partitionNet.String()
}

func nextPowerOf2(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++

	return v
}
