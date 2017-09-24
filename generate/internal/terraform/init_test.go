package terraform_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTerraform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "generate/internal/terraform")
}
