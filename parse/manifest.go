package parse

type Manifest struct {
	Provider *ManifestProvider `yaml:"provider"`
	Network  ManifestNetwork   `yaml:"network"`
}
