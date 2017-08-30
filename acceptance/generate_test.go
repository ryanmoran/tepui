package acceptance

import (
	"io/ioutil"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("generate", func() {
	Describe("GCP", func() {
		It("generates a template with a provider", func() {
			command := exec.Command(pathToMain,
				"--manifest", "fixtures/manifests/gcp.yml")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			contents, err := ioutil.ReadFile("fixtures/templates/gcp.json")
			Expect(err).NotTo(HaveOccurred())
			Expect(session.Out.Contents()).Should(MatchJSON(contents))
		})
	})

	Describe("AWS", func() {
		It("generates a template with a provider", func() {
			command := exec.Command(pathToMain,
				"--manifest", "fixtures/manifests/aws.yml")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			contents, err := ioutil.ReadFile("fixtures/templates/aws.json")
			Expect(err).NotTo(HaveOccurred())
			Expect(session.Out.Contents()).Should(MatchJSON(contents))
		})
	})

	Describe("Azure", func() {
		PIt("generates a template with a provider", func() {
			command := exec.Command(pathToMain,
				"--manifest", "fixtures/manifests/azure.yml")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))

			contents, err := ioutil.ReadFile("fixtures/templates/azure.json")
			Expect(err).NotTo(HaveOccurred())
			Expect(session.Out.Contents()).Should(MatchJSON(contents))
		})
	})
})
