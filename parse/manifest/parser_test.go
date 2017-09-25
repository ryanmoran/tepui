package manifest_test

import (
	"github.com/ryanmoran/tepui/parse/manifest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Parse", func() {
		var parser manifest.Parser

		BeforeEach(func() {
			parser = manifest.NewParser()
		})

		It("parses a manifest from a file path", func() {
			m, err := parser.Parse("fixtures/manifest.yml")
			Expect(err).NotTo(HaveOccurred())
			Expect(m).To(Equal(manifest.Manifest{
				Name: "some-environment",
				Networks: []manifest.Network{
					{
						Name: "some-network",
						CIDR: "1.2.3.4/5",
						Subnets: []manifest.Subnet{
							{
								Name: "some-subnet",
								CIDR: "2.3.4.5/6",
							},
						},
					},
				},
				LoadBalancers: []manifest.LoadBalancer{
					{
						Name:  "some-lb",
						Ports: []int{1234},
						Zones: []string{"some-zone"},
					},
				},
			}))
		})

		Describe("error cases", func() {
			Context("when the manifest file cannot be found", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/notfound.yml")
					Expect(err).To(MatchError("open fixtures/notfound.yml: no such file or directory"))
				})
			})

			Context("when the manifest cannot be unmarshaled", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/malformed.yml")
					Expect(err).To(MatchError(ContainSubstring("yaml: could not find expected directive name")))
				})
			})
		})
	})
})
