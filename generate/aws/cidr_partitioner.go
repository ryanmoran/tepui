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

	//We get the number of leading ones and total bits in the mask of the IP network from the given CIDR.
	//
	//Example:
	//  CIDR = 10.64.0.0/10
	//  ones = 8
	//  bits = 32
	ones, bits := ipNet.Mask.Size()

	//We calculate the total number of available IPs in the network.
	//We determine the next power of two for the number of parts we will split the network into.
	//
	//Example:
	//  2 ^ (bits - ones) = 16777216
	//  total number of available IPs / number of parts = 4194304
	partition := int(math.Pow(2, float64(bits-ones))) / nextPowerOf2(p.Parts)

	//We determine that `ones` is actually the total bits minus the bits for this parition
	//
	//Example:
	//  log partition = 22
	//  ones = 10
	ones = bits - int(math.Log2(float64(partition)))

	//We calculate the number of IPs in this subnet.
	//Example:
	//  n = 167772160
	//  n + (index times 2, 22 times) = 171966464
	n := binary.BigEndian.Uint32(ipNet.IP)
	n += uint32(index) << uint(bits-ones)

	// This creates a []byte for the subnet.
	subnetIP := make(net.IP, len(ipNet.IP))
	binary.BigEndian.PutUint32(subnetIP, n)

	// This creates a net.IPNet type for the subnet with the updated mask.
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
