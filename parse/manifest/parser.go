package manifest

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(path string) (Manifest, error) {
	var manifest Manifest

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return manifest, err
	}

	err = yaml.Unmarshal(contents, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}
