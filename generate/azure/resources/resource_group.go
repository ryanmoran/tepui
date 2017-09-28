package resources

type AzurermResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func NewAzurermResourceGroup(name, location string) AzurermResourceGroup {
	return AzurermResourceGroup{
		Name:     name,
		Location: location,
	}
}
