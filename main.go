package main

import (
	"encoding/json"
	"flag"
	"fmt"

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

	//generator := generate.TemplateGenerator(manifest)
	//template, err := generator.Generate()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(template)

	var template Template
	template.Provider.Google.Credentials = manifest.Provider.GCP.Credentials
	template.Provider.Google.Project = manifest.Provider.GCP.Project
	template.Provider.Google.Region = manifest.Provider.GCP.Region
	template.Resource.GoogleComputeNetwork.Network.Name = manifest.Network.Name

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
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
