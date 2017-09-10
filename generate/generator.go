package generate

import (
	"github.com/pivotal-cf/tepui/parse"
	"github.com/pivotal-cf/tepui/parse/provider"
)

type Generator interface {
	Generate(provider.Provider, parse.Manifest) (template string, err error)
}
