package gcp

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate/gcp/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse"
	"github.com/pivotal-cf/tepui/parse/provider"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(prov provider.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate(Provider{
		Credentials: prov.GCP.Credentials,
		Project:     prov.GCP.Project,
		Region:      prov.GCP.Region,
	})

	for _, network := range manifest.Networks {
		networkResource := terraform.NamedResource{
			Name: network.Name,
			Resource: resources.GoogleComputeNetwork{
				Name: network.Name,
			},
		}

		template.Resources = append(template.Resources, networkResource)

		for _, subnet := range network.Subnets {
			subnetResource := terraform.NamedResource{
				Name: subnet.Name,
				Resource: resources.GoogleComputeSubnetwork{
					Name:        subnet.Name,
					IPCIDRRange: subnet.CIDR,
					Network:     networkResource.SelfLink(),
				},
			}

			template.Resources = append(template.Resources, subnetResource)
		}
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
