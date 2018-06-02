package generator_test

import (
	"io/ioutil"

	"github.com/calebwashburn/tile-config-template-generator/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Product Properties", func() {
	Context("NewProductProperties", func() {
		It("Should return new required product properties", func() {
			fileData, err := ioutil.ReadFile("fixtures/p_healthwatch.yml")
			Expect(err).ShouldNot(HaveOccurred())
			expected, err := ioutil.ReadFile("fixtures/healthwatch-required.yml")
			Expect(err).ShouldNot(HaveOccurred())
			metadata, err := generator.NewMetadata(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			productProperties, err := generator.NewProductProperties(metadata)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productProperties).ShouldNot(BeNil())
			yml, err := yaml.Marshal(productProperties)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(yml).Should(MatchYAML(string(expected)))
		})

		It("Should return new required product properties", func() {
			fileData, err := ioutil.ReadFile("fixtures/pas.yml")
			Expect(err).ShouldNot(HaveOccurred())
			expected, err := ioutil.ReadFile("fixtures/pas-required.yml")
			Expect(err).ShouldNot(HaveOccurred())
			metadata, err := generator.NewMetadata(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			productProperties, err := generator.NewProductProperties(metadata)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productProperties).ShouldNot(BeNil())
			yml, err := yaml.Marshal(productProperties)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(yml).Should(MatchYAML(string(expected)))
		})
	})
})
