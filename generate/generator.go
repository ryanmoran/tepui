package generate

import "github.com/pivotal-cf/tepui/parse"

type Generator interface {
	Generate(parse.Provider, parse.Manifest) (template string, err error)
}
