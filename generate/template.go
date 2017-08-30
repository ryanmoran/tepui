package generate

type Template struct {
	Providers *TemplateProviderCollection `json:"provider"`
	Resources *TemplateResourceCollection `json:"resource"`
}

func NewTemplate() Template {
	return Template{
		Providers: &TemplateProviderCollection{},
		Resources: &TemplateResourceCollection{},
	}
}
