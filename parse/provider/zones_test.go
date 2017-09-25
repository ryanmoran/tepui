package provider_test

import (
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zones", func() {
	Describe("Find", func() {
		var zones provider.Zones

		BeforeEach(func() {
			zones = provider.Zones{
				{
					Alias: "zone-0",
					Name:  "some-zone-0",
				},
				{
					Alias: "zone-1",
					Name:  "some-zone-1",
				},
				{
					Alias: "zone-2",
					Name:  "some-zone-2",
				},
			}
		})

		It("returns the zone name matching the given alias", func() {
			zone, ok := zones.Find("zone-1")
			Expect(zone).To(Equal("some-zone-1"))
			Expect(ok).To(BeTrue())
		})

		It("returns an empty string and false for a missing alias", func() {
			zone, ok := zones.Find("zone-6")
			Expect(zone).To(Equal(""))
			Expect(ok).To(BeFalse())
		})
	})
})
