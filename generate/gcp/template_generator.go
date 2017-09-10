package gcp

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate/gcp/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate(Provider{
		Credentials: provider.GCP.Credentials,
		Project:     provider.GCP.Project,
		Region:      provider.GCP.Region,
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
