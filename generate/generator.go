package generate

import (
	"github.com/pivotal-cf/tepui/parse/manifest"
	"github.com/pivotal-cf/tepui/parse/provider"
)

type Generator interface {
	Generate(provider.Provider, manifest.Manifest) (template string, err error)
}
