package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pivotal-cf/tepui/generate"
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

	prov, err := provider.NewParser().Parse(providerPath)
	if err != nil {
		log.Fatalln(err)
	}

	manifest, err := manifest.NewParser().Parse(manifestPath)
	if err != nil {
		log.Fatalln(err)
	}

	var templateGenerator generate.Generator

	switch prov.Type {
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

	template, err := templateGenerator.Generate(prov, manifest)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(template)
}
