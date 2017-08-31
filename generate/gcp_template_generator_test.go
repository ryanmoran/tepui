package generate_test

import (
	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GCPTemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			provider := parse.Provider{
				Type: "gcp",
				GCP: parse.ProviderGCP{
					Credentials: "some-credentials",
					Project:     "some-project",
					Region:      "some-region",
				},
			}
			manifest := parse.Manifest{
				Name: "some-environment",
				Networks: []parse.ManifestNetwork{
					{
						Name: "some-network",
						CIDR: "1.2.3.4/5",
					},
				},
			}

			template, err := generate.NewGCPTemplateGenerator().Generate(provider, manifest)
			Expect(err).NotTo(HaveOccurred())
			Expect(template).To(MatchJSON(`{
				"provider": {
					"google": {
						"credentials": "some-credentials",
						"project": "some-project",
						"region": "some-region"
					}
				},
				"resource": {
					"google_compute_network": {
						"some-network": {
							"name": "some-network"
						}
					}
				}
			}`))
		})
	})
})
