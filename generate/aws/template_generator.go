package aws

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := generate.NewTemplate()

	template.Providers.Add(Provider{
		AccessKey: provider.AWS.AccessKey,
		SecretKey: provider.AWS.SecretKey,
		Region:    provider.AWS.Region,
	})

	for _, network := range manifest.Networks {
		template.Resources.Add(network.Name, VPC{
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
