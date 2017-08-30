package parse

type ManifestEnvironment struct {
	Name     string                       `yaml:"name"`
	Networks []ManifestEnvironmentNetwork `yaml:"networks"`
}
