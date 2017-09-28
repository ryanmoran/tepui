package azure_test

import (
	"github.com/ryanmoran/tepui/generate/azure"
	"github.com/ryanmoran/tepui/generate/azure/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NetworkResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a network", func() {
			generator := azure.NewNetworkResourceGenerator()

			resourceGroup := terraform.NamedResource{
				Name: "some-resource-group",
				Resource: resources.AzurermResourceGroup{
					Name: "some-resource-group",
				},
			}

			network := manifest.Network{
				Name: "some-network",
				CIDR: "1.2.3.4/5",
				Subnets: []manifest.Subnet{
					{
						Name: "some-subnet",
						CIDR: "6.7.8.9/10",
					},
				},
			}

			Expect(generator.Generate(resourceGroup, "some-region", network)).To(ConsistOf(terraform.Resources{
				{
					Name: "some-network",
					Resource: resources.AzurermVirtualNetwork{
						Name:              "some-network",
						ResourceGroupName: "${azurerm_resource_group.some-resource-group.name}",
						AddressSpace:      []string{"1.2.3.4/5"},
						Location:          "some-region",
					},
				},
				{
					Name: "some-subnet",
					Resource: resources.AzurermSubnet{
						Name:               "some-subnet",
						ResourceGroupName:  "${azurerm_resource_group.some-resource-group.name}",
						VirtualNetworkName: "${azurerm_virtual_network.some-network.name}",
						AddressPrefix:      "6.7.8.9/10",
					},
				},
			}))
		})
	})
})
