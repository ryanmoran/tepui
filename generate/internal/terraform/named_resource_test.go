package terraform_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
)

type SomeCloudResource struct{}

var _ = Describe("NamedResource", func() {
	var namedResource terraform.NamedResource

	BeforeEach(func() {
		namedResource = terraform.NamedResource{
			Name:     "some-resource",
			Resource: SomeCloudResource{},
		}
	})

	Describe("Type", func() {
		It("returns a snake-cased string of the struct type of the Resource", func() {
			Expect(namedResource.Type()).To(Equal("some_cloud_resource"))
		})
	})

	Describe("Attribute", func() {
		It("returns the terraform attribute link for this named resource", func() {
			Expect(namedResource.Attribute("something")).To(Equal("${some_cloud_resource.some-resource.something}"))
		})
	})
})
