package generate_test

import (
	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			manifest := parse.Manifest{
				Provider: &parse.ManifestProvider{
					Type: "gcp",
					GCP: parse.ManifestProviderGCP{
						Credentials: "some-credentials",
						Project:     "some-project",
						Region:      "some-region",
					},
				},
				Network: parse.ManifestNetwork{
					Name: "some-network",
				},
			}

			generator := generate.NewTemplateGenerator()
			template, err := generator.Generate(manifest)
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
						"network": {
							"name": "some-network"
						}
					}
				}
			}`))
		})
	})
})
