package aws_test

import (
	"github.com/ryanmoran/tepui/generate/aws"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CIDRPartitioner", func() {
	Describe("Partition", func() {
		It("returns the CIDR partion for the given index", func() {
			partitioner := aws.NewCIDRPartitioner("10.0.0.0/8", 3)
			Expect(partitioner.Partition(0)).To(Equal("10.0.0.0/10"))
			Expect(partitioner.Partition(1)).To(Equal("10.64.0.0/10"))
			Expect(partitioner.Partition(2)).To(Equal("10.128.0.0/10"))

			partitioner = aws.NewCIDRPartitioner("192.168.0.0/16", 8)
			Expect(partitioner.Partition(0)).To(Equal("192.168.0.0/19"))
			Expect(partitioner.Partition(1)).To(Equal("192.168.32.0/19"))
			Expect(partitioner.Partition(2)).To(Equal("192.168.64.0/19"))
			Expect(partitioner.Partition(3)).To(Equal("192.168.96.0/19"))
			Expect(partitioner.Partition(4)).To(Equal("192.168.128.0/19"))
			Expect(partitioner.Partition(5)).To(Equal("192.168.160.0/19"))
			Expect(partitioner.Partition(6)).To(Equal("192.168.192.0/19"))
			Expect(partitioner.Partition(7)).To(Equal("192.168.224.0/19"))

			partitioner = aws.NewCIDRPartitioner("8.8.8.0/24", 10)
			Expect(partitioner.Partition(0)).To(Equal("8.8.8.0/28"))
			Expect(partitioner.Partition(1)).To(Equal("8.8.8.16/28"))
			Expect(partitioner.Partition(2)).To(Equal("8.8.8.32/28"))
			Expect(partitioner.Partition(3)).To(Equal("8.8.8.48/28"))
			Expect(partitioner.Partition(4)).To(Equal("8.8.8.64/28"))
			Expect(partitioner.Partition(5)).To(Equal("8.8.8.80/28"))
			Expect(partitioner.Partition(6)).To(Equal("8.8.8.96/28"))
			Expect(partitioner.Partition(7)).To(Equal("8.8.8.112/28"))
			Expect(partitioner.Partition(8)).To(Equal("8.8.8.128/28"))
			Expect(partitioner.Partition(9)).To(Equal("8.8.8.144/28"))
		})
	})
})
