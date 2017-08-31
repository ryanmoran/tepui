package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type AWSTemplateGenerator struct{}

func NewAWSTemplateGenerator() AWSTemplateGenerator {
	return AWSTemplateGenerator{}
}

func (g AWSTemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate()

	template.Providers.Add(TemplateProviderAWS{
		AccessKey: provider.AWS.AccessKey,
		SecretKey: provider.AWS.SecretKey,
		Region:    provider.AWS.Region,
	})

	for _, network := range manifest.Networks {
		template.Resources.Add(network.Name, TemplateResourceAWSVPC{
			CIDRBlock: network.CIDR,
			Tags: map[string]string{
				"name": network.Name,
			},
		})
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
