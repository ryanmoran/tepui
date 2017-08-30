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
