package generate_test

import (
	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		Context("for GCP", func() {
			It("generates a template from the given manifest", func() {
				manifest := parse.Manifest{
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
				}

				generator := generate.NewTemplateGenerator()
				template, err := generator.Generate(manifest)
				Expect(err).NotTo(HaveOccurred())
				Expect(template).To(MatchJSON(`{
					"provider": {
						"google": {
							"credentials": "some-credentials",
							"project": "some-project",
							"region": "some-region"
						}
					},
					"resource": {
						"google_compute_network": {
							"some-network": {
								"name": "some-network"
							}
						}
					}
				}`))
			})
		})

		Context("for AWS", func() {
			It("generates a template from the given manifest", func() {
				manifest := parse.Manifest{
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
				}

				generator := generate.NewTemplateGenerator()
				template, err := generator.Generate(manifest)
				Expect(err).NotTo(HaveOccurred())
				Expect(template).To(MatchJSON(`{
					"provider": {
						"aws": {
							"access_key": "some-access-key",
							"secret_key": "some-secret-key",
							"region": "some-region"
						}
					},
					"resource": {
						"aws_vpc": {
							"some-network": {
								"cidr_block": "1.2.3.4/5",
								"tags": {
									"name": "some-network"
								}
							}
						}
					}
				}`))
			})
		})

		Context("for Azure", func() {
			It("generates a template from the given manifest", func() {
				manifest := parse.Manifest{
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
				}

				generator := generate.NewTemplateGenerator()
				template, err := generator.Generate(manifest)
				Expect(err).NotTo(HaveOccurred())
				Expect(template).To(MatchJSON(`{
					"provider": {
						"azurerm": {
							"subscription_id": "some-subscription-id",
							"client_id": "some-client-id",
							"client_secret": "some-client-secret",
							"tenant_id": "some-tenant-id"
						}
					},
					"resource": {
						"azurerm_resource_group": {
							"resource_group": {
								"name": "some-environment",
								"location": "some-region"
							}
						},
						"azurerm_virtual_network": {
							"some-network": {
								"name": "some-network",
								"resource_group_name": "${azurerm_resource_group.resource_group.name}",
								"address_space": ["1.2.3.4/5"],
								"location": "some-region"
							}
						}
					}
				}`))
			})
		})
	})
})
