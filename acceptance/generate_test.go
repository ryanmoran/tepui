package acceptance

import (
	"io/ioutil"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("generate", func() {
	var templateFixture string

	BeforeEach(func() {
		contents, err := ioutil.ReadFile("fixtures/template.json")
		Expect(err).NotTo(HaveOccurred())

		templateFixture = string(contents)
	})

	It("generates a template with a provider", func() {
		command := exec.Command(pathToMain,
			"--manifest", "fixtures/environment.yml")
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(session.Out.Contents()).Should(MatchJSON(templateFixture))
	})
})
