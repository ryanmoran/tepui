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

	manifest, err := manifest.NewManifestParser().Parse(manifestPath)
	if err != nil {
		log.Fatalln(err)
	}

	var generator generate.Generator

	switch prov.Type {
	case "gcp":
		generator = gcp.NewTemplateGenerator()
	case "aws":
		generator = aws.NewTemplateGenerator()
	case "azure":
		generator = azure.NewTemplateGenerator()
	}

	template, err := generator.Generate(prov, manifest)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(template)
}
