package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pivotal-cf/tepui/generate/aws"
	"github.com/pivotal-cf/tepui/generate/azure"
	"github.com/pivotal-cf/tepui/generate/gcp"
	"github.com/pivotal-cf/tepui/parse/manifest"
	"github.com/pivotal-cf/tepui/parse/provider"
)

func main() {
	var (
		manifestPath string
		providerPath string
	)

	flag.StringVar(&providerPath, "provider", "", "path to provider")
	flag.StringVar(&manifestPath, "manifest", "", "path to manifest")
	flag.Parse()

	p, err := provider.NewParser().Parse(providerPath)
	if err != nil {
		log.Fatalln(err)
	}

	m, err := manifest.NewParser().Parse(manifestPath)
	if err != nil {
		log.Fatalln(err)
	}

	var templateGenerator interface {
		Generate(provider.Provider, manifest.Manifest) (string, error)
	}

	switch p.Type {

	case "aws":
		networksGenerator := aws.NewNetworkResourceGenerator()
		templateGenerator = aws.NewTemplateGenerator(networksGenerator)

	case "azure":
		networksGenerator := azure.NewNetworkResourceGenerator()
		templateGenerator = azure.NewTemplateGenerator(networksGenerator)

	case "gcp":
		networksGenerator := gcp.NewNetworkResourceGenerator()
		templateGenerator = gcp.NewTemplateGenerator(networksGenerator)

	}

	template, err := templateGenerator.Generate(p, m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(template)
}
