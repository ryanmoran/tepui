package parse_test

import (
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProviderParser", func() {
	Describe("Parse", func() {
		var parser parse.ProviderParser

		BeforeEach(func() {
			parser = parse.NewProviderParser()
		})

		Context("for GCP", func() {
			It("parses a provider from a file path", func() {
				provider, err := parser.Parse("fixtures/providers/gcp.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(provider).To(Equal(parse.Provider{
					Type: "gcp",
					GCP: parse.ProviderGCP{
						Credentials: "some-credentials",
						Project:     "some-project",
						Region:      "some-region",
					},
				}))
			})
		})

		Context("for AWS", func() {
			It("parses a provider from a file path", func() {
				provider, err := parser.Parse("fixtures/providers/aws.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(provider).To(Equal(parse.Provider{
					Type: "aws",
					AWS: parse.ProviderAWS{
						AccessKey: "some-access-key",
						SecretKey: "some-secret-key",
						Region:    "some-region",
					},
				}))
			})
		})

		Context("for Azure", func() {
			It("parses a provider from a file path", func() {
				provider, err := parser.Parse("fixtures/providers/azure.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(provider).To(Equal(parse.Provider{
					Type: "azure",
					Azure: parse.ProviderAzure{
						SubscriptionID: "some-subscription-id",
						ClientID:       "some-client-id",
						ClientSecret:   "some-client-secret",
						TenantID:       "some-tenant-id",
						Region:         "some-region",
					},
				}))
			})
		})

		Describe("error cases", func() {
			Context("when the provider file cannot be found", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/notfound.yml")
					Expect(err).To(MatchError("open fixtures/notfound.yml: no such file or directory"))
				})
			})

			Context("when the provider cannot be unmarshaled", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/providers/malformed.yml")
					Expect(err).To(MatchError(ContainSubstring("found character that cannot start any token")))
				})
			})

			Context("when an unknown provider type is specified", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/providers/unknown.yml")
					Expect(err).To(MatchError(ContainSubstring("unknown provider type: \"banana\"")))
				})
			})
		})
	})
})
