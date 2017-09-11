package manifest

type Manifest struct {
	Name     string            `yaml:"name"`
	Networks []ManifestNetwork `yaml:"networks"`
}
