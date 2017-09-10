package terraform

import (
	"fmt"
	"reflect"
	"unicode"
)

type NamedResource struct {
	Name     string
	Resource Resource
}

func (nr NamedResource) Type() string {
	return camelToSnake(reflect.TypeOf(nr.Resource).Name())
}

func (nr NamedResource) Attribute(attr string) string {
	return fmt.Sprintf("${%s.%s.%s}", nr.Type(), nr.Name, attr)
}

func (nr NamedResource) SelfLink() string {
	return nr.Attribute("self_link")
}

func camelToSnake(input string) string {
	var output string

	for index, r := range input {
		if unicode.IsUpper(r) {
			if index != 0 {
				output += "_"
			}
			output += string(unicode.ToLower(r))
		} else {
			output += string(r)
		}
	}

	return output
}
