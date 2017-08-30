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

	output, err := json.Marshal(template)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func NewTemplate() Template {
	return Template{
		Providers: &TemplateProviderCollection{},
		Resources: &TemplateResourceCollection{},
	}
}

type Template struct {
	Providers *TemplateProviderCollection `json:"provider"`
	Resources *TemplateResourceCollection `json:"resource"`
}

type provider interface {
	_provider()
}

type TemplateProviderCollection struct {
	providers map[string]provider
}

func (tpc *TemplateProviderCollection) Add(p provider) {
	if tpc.providers == nil {
		tpc.providers = make(map[string]provider)
	}

	switch p.(type) {
	case TemplateProviderGoogle:
		tpc.providers["google"] = p
	case TemplateProviderAWS:
		tpc.providers["aws"] = p
	}
}

func (tpc TemplateProviderCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(tpc.providers)
}

type TemplateProviderGoogle struct {
	Credentials string `json:"credentials"`
	Project     string `json:"project"`
	Region      string `json:"region"`
}

func (tpg TemplateProviderGoogle) _provider() {}

type TemplateProviderAWS struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

func (tpa TemplateProviderAWS) _provider() {}

type resource interface {
	resourceType() string
}

type TemplateResourceCollection struct {
	resources map[string]map[string]resource
}

func (trc *TemplateResourceCollection) Add(name string, r resource) {
	if trc.resources == nil {
		trc.resources = make(map[string]map[string]resource)
	}

	if trc.resources[r.resourceType()] == nil {
		trc.resources[r.resourceType()] = make(map[string]resource)
	}

	trc.resources[r.resourceType()][name] = r
}

func (trc TemplateResourceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(trc.resources)
}

type TemplateResourceGoogleComputeNetwork struct {
	Name string `json:"name"`
}

func (trgcn TemplateResourceGoogleComputeNetwork) resourceType() string {
	return "google_compute_network"
}

type TemplateResourceAWSVPC struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func (trav TemplateResourceAWSVPC) resourceType() string {
	return "aws_vpc"
}
