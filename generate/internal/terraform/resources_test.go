package terraform_test

import (
	"github.com/pivotal-cf/tepui/generate/internal/terraform"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type SomeAwsType struct {
	SomeProperty string `json:"some_property"`
}

var _ = Describe("Resources", func() {
	Describe("MarshalJSON", func() {
		It("returns a JSON representation of the resource collection", func() {
			resources := terraform.Resources{
				{
					Name: "some_resource_name",
					Resource: SomeAwsType{
						SomeProperty: "banana",
					},
				},
			}
			Expect(resources.MarshalJSON()).To(MatchJSON(`{
				"some_aws_type": {
					"some_resource_name": {
						"some_property": "banana"
					}
				}
			}`))
		})
	})
})
