package resources

type AwsSecurityGroup struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
}

func NewAwsSecurityGroup(name string) AwsSecurityGroup {
	return AwsSecurityGroup{
		Name: name,
		Tags: map[string]string{
			"Name": name,
		},
	}
}
