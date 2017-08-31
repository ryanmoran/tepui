package azure_test

import (
	"github.com/pivotal-cf/tepui/generate/azure"
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			provider := parse.Provider{
				Type: "azure",
				Azure: parse.ProviderAzure{
					SubscriptionID: "some-subscription-id",
					ClientID:       "some-client-id",
					ClientSecret:   "some-client-secret",
					TenantID:       "some-tenant-id",
					Region:         "some-region",
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

			generator := azure.NewTemplateGenerator()
			template, err := generator.Generate(provider, manifest)
			Expect(err).NotTo(HaveOccurred())
			Expect(template).To(MatchJSON(`{
				"provider": {
					"azurerm": {
						"subscription_id": "some-subscription-id",
						"client_id": "some-client-id",
						"client_secret": "some-client-secret",
						"tenant_id": "some-tenant-id"
					}
				},
				"resource": {
					"azurerm_resource_group": {
						"resource_group": {
							"name": "some-environment",
							"location": "some-region"
						}
					},
					"azurerm_virtual_network": {
						"some-network": {
							"name": "some-network",
							"resource_group_name": "${azurerm_resource_group.resource_group.name}",
							"address_space": ["1.2.3.4/5"],
							"location": "some-region"
						}
					}
				}
			}`))
		})
	})
})
