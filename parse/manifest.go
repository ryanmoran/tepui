package parse

type Manifest struct {
	Provider    *ManifestProvider   `yaml:"provider"`
	Environment ManifestEnvironment `yaml:"environment"`
}
