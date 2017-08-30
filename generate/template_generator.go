package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (tg TemplateGenerator) Generate(manifest parse.Manifest) (string, error) {
	template := NewTemplate()
	switch manifest.Provider.Type {
	case "gcp":
		template.Providers.Add(TemplateProviderGoogle{
			Credentials: manifest.Provider.GCP.Credentials,
			Project:     manifest.Provider.GCP.Project,
			Region:      manifest.Provider.GCP.Region,
		})

		template.Resources.Add("network", TemplateResourceGoogleComputeNetwork{
			Name: manifest.Network.Name,
		})
	case "aws":
		template.Providers.Add(TemplateProviderAWS{
			AccessKey: manifest.Provider.AWS.AccessKey,
			SecretKey: manifest.Provider.AWS.SecretKey,
			Region:    manifest.Provider.AWS.Region,
		})

		template.Resources.Add("network", TemplateResourceAWSVPC{
			CIDRBlock: manifest.Network.CIDR,
			Tags: map[string]string{
				"name": manifest.Network.Name,
			},
		})
	default:
		panic("unknown provider")
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
