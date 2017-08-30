package main

import (
	"flag"
	"fmt"

	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"
)

func main() {
	var manifestPath string
	flag.StringVar(&manifestPath, "manifest", "", "path to manifest")
	flag.Parse()

	parser := parse.NewManifestParser()
	manifest, err := parser.Parse(manifestPath)
	if err != nil {
		panic(err)
	}

	generator := generate.NewTemplateGenerator()
	template, err := generator.Generate(manifest)
	if err != nil {
		panic(err)
	}

	fmt.Println(template)
}

type Manifest struct {
	Provider struct {
		Type   string `yaml:"type"`
		Params struct {
			Credentials string `yaml:"credentials"`
			Project     string `yaml:"project"`
			Region      string `yaml:"region"`
		} `yaml:"params"`
	} `yaml:"provider"`
	Network struct {
		Name string `yaml:"name"`
	} `yaml:"network"`
}

type Template struct {
	Provider struct {
		Google struct {
			Credentials string `json:"credentials"`
			Project     string `json:"project"`
			Region      string `json:"region"`
		} `json:"google"`
	} `json:"provider"`
	Resource struct {
		GoogleComputeNetwork struct {
			Network struct {
				Name string `json:"name"`
			} `json:"network"`
		} `json:"google_compute_network"`
	} `json:"resource"`
}
