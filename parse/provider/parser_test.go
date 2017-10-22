package provider_test

import (
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Parse", func() {
		var parser provider.Parser

		BeforeEach(func() {
			parser = provider.NewParser()
		})

		Context("for GCP", func() {
			It("parses a provider from a file path", func() {
				gcp, err := parser.Parse("fixtures/gcp.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(gcp).To(Equal(provider.Provider{
					Type: "gcp",
					GCP: provider.GCP{
						Credentials: "some-credentials",
						Project:     "some-project",
						Region:      "some-region",
						Zones: []provider.Zone{
							{
								Alias: "zone-1",
								Name:  "name-1",
							},
						},
					},
				}))
			})
		})

		Context("for AWS", func() {
			It("parses a provider from a file path", func() {
				aws, err := parser.Parse("fixtures/aws.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(aws).To(Equal(provider.Provider{
					Type: "aws",
					AWS: provider.AWS{
						AccessKey: "some-access-key",
						SecretKey: "some-secret-key",
						Region:    "some-region",
						Zones: []provider.Zone{
							{
								Alias: "zone-1",
								Name:  "name-1",
							},
						},
					},
				}))
			})
		})

		Context("for Azure", func() {
			It("parses a provider from a file path", func() {
				azure, err := parser.Parse("fixtures/azure.yml")
				Expect(err).NotTo(HaveOccurred())
				Expect(azure).To(Equal(provider.Provider{
					Type: "azure",
					Azure: provider.Azure{
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
					_, err := parser.Parse("fixtures/malformed.yml")
					Expect(err).To(MatchError(ContainSubstring("found character that cannot start any token")))
				})
			})

			Context("when an unknown provider type is specified", func() {
				It("returns an error", func() {
					_, err := parser.Parse("fixtures/unknown.yml")
					Expect(err).To(MatchError(ContainSubstring("unknown provider type: \"banana\"")))
				})
			})
		})
	})
})
