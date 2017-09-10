package aws

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate/aws/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate(Provider{
		AccessKey: provider.AWS.AccessKey,
		SecretKey: provider.AWS.SecretKey,
		Region:    provider.AWS.Region,
	})

	for _, network := range manifest.Networks {
		networkResource := terraform.NamedResource{
			Name: network.Name,
			Resource: resources.AwsVpc{
				CIDRBlock: network.CIDR,
				Tags: map[string]string{
					"Name":        network.Name,
					"Environment": manifest.Name,
				},
			},
		}

		template.Resources = append(template.Resources, networkResource)

		for _, subnet := range network.Subnets {
			subnetResource := terraform.NamedResource{
				Name: subnet.Name,
				Resource: resources.AwsSubnet{
					VPCID:     networkResource.Attribute("id"),
					CIDRBlock: subnet.CIDR,
					Tags: map[string]string{
						"Name":        subnet.Name,
						"Environment": manifest.Name,
					},
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
