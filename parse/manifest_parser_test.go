package parse_test

import (
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ManifestParser", func() {
	Describe("Parse", func() {
		var parser parse.ManifestParser

		BeforeEach(func() {
			parser = parse.NewManifestParser()
		})

		Context("for GCP", func() {
			It("parses a manifest from a file path", func() {
				manifest, err := parser.Parse("fixtures/gcp.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(manifest).To(Equal(parse.Manifest{
					Provider: &parse.ManifestProvider{
						Type: "gcp",
						GCP: parse.ManifestProviderGCP{
							Credentials: "some-credentials",
							Project:     "some-project",
							Region:      "some-region",
						},
					},
					Environment: parse.ManifestEnvironment{
						Name: "some-environment",
						Networks: []parse.ManifestEnvironmentNetwork{
							{
								Name: "some-network",
								CIDR: "1.2.3.4/5",
							},
						},
					},
				}))
			})
		})

		Context("for AWS", func() {
			It("parses a manifest from a file path", func() {
				manifest, err := parser.Parse("fixtures/aws.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(manifest).To(Equal(parse.Manifest{
					Provider: &parse.ManifestProvider{
						Type: "aws",
						AWS: parse.ManifestProviderAWS{
							AccessKey: "some-access-key",
							SecretKey: "some-secret-key",
							Region:    "some-region",
						},
					},
					Environment: parse.ManifestEnvironment{
						Name: "some-environment",
						Networks: []parse.ManifestEnvironmentNetwork{
							{
								Name: "some-network",
								CIDR: "1.2.3.4/5",
							},
						},
					},
				}))
			})
		})

		Context("for Azure", func() {
			It("parses a manifest from a file path", func() {
				manifest, err := parser.Parse("fixtures/azure.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(manifest).To(Equal(parse.Manifest{
					Provider: &parse.ManifestProvider{
						Type: "azure",
						Azure: parse.ManifestProviderAzure{
							SubscriptionID: "some-subscription-id",
							ClientID:       "some-client-id",
							ClientSecret:   "some-client-secret",
							TenantID:       "some-tenant-id",
							Region:         "some-region",
						},
					},
					Environment: parse.ManifestEnvironment{
						Name: "some-environment",
						Networks: []parse.ManifestEnvironmentNetwork{
							{
								Name: "some-network",
								CIDR: "1.2.3.4/5",
							},
						},
					},
				}))
			})
		})

		Describe("error cases", func() {
			Context("when the manifest file cannot be found", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/notfound.yml")
					Expect(err).To(MatchError("open fixtures/notfound.yml: no such file or directory"))
				})
			})

			Context("when the manifest cannot be unmarshaled", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/malformed.yml")
					Expect(err).To(MatchError(ContainSubstring("yaml: could not find expected directive name")))
				})
			})

			Context("when the manifest provider key cannot be unmarshalled", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/malformed-provider.yml")
					Expect(err).To(MatchError(ContainSubstring("cannot unmarshal")))
				})
			})

			Context("when an unknown provider type is specified", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/unknown-provider.yml")
					Expect(err).To(MatchError(ContainSubstring("unknown provider type: \"banana\"")))
				})
			})
		})
	})
})
