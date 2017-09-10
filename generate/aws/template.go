package aws

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		VPCs Resources `json:"aws_vpc"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
