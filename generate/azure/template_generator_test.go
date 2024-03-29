package azure_test

import (
	"github.com/ryanmoran/tepui/generate/azure"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			prov := provider.Provider{
				Type: "azure",
				Azure: provider.Azure{
					SubscriptionID: "some-subscription-id",
					ClientID:       "some-client-id",
					ClientSecret:   "some-client-secret",
					TenantID:       "some-tenant-id",
					Region:         "some-region",
				},
			}
			manifest := manifest.Manifest{
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
			}

			generator := azure.NewTemplateGenerator(azure.NewNetworkResourceGenerator())
			template, err := generator.Generate(prov, manifest)
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
					},
					"azurerm_subnet": {
						"some-subnet": {
							"name": "some-subnet",
							"resource_group_name": "${azurerm_resource_group.resource_group.name}",
							"virtual_network_name": "${azurerm_virtual_network.some-network.name}",
							"address_prefix": "2.3.4.5/6"
						}
					}
				}
			}`))
		})
	})
})
