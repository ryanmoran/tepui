package main

import (
	"flag"
	"fmt"

	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/generate/aws"
	"github.com/pivotal-cf/tepui/generate/azure"
	"github.com/pivotal-cf/tepui/generate/gcp"
	"github.com/pivotal-cf/tepui/parse"
)

func main() {
	var (
		manifestPath string
		providerPath string
	)

	flag.StringVar(&providerPath, "provider", "", "path to provider")
	flag.StringVar(&manifestPath, "manifest", "", "path to manifest")
	flag.Parse()

	provider, err := parse.NewProviderParser().Parse(providerPath)
	if err != nil {
		panic(err)
	}

	manifest, err := parse.NewManifestParser().Parse(manifestPath)
	if err != nil {
		panic(err)
	}

	var generator generate.Generator

	switch provider.Type {
	case "gcp":
		generator = gcp.NewTemplateGenerator()
	case "aws":
		generator = aws.NewTemplateGenerator()
	case "azure":
		generator = azure.NewTemplateGenerator()
	}

	template, err := generator.Generate(provider, manifest)
	if err != nil {
		panic(err)
	}

	fmt.Println(template)
}
