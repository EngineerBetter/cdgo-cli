package dir_test

import (
	"github.com/EngineerBetter/cdgo/dir"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"path/filepath"
)

var _ = Describe("finder", func() {
	var (
		cwd string
	)

	BeforeEach(func() {
		cwd, _ = os.Getwd()
	})

	Describe("find", func() {
		Context("when the starting directory does not exist", func() {
			It("returns an error", func() {
				startDir := filepath.Join(cwd, "not-exist")
				_, err := dir.Find("needle", startDir)
				Ω(err).ShouldNot(BeNil(), "Error should be returned for non-existant directory")
			})
		})

		Context("when the starting directory exists", func() {
			Context("and the target directory exists", func() {
				It("returns the absolute path to the highest matching directory", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					result, err := dir.Find("root", startDir)
					Ω(err).Should(BeNil())
					Ω(result).Should(Equal(cwd + "/test-fixtures/root"))
				})
			})

			Context("but the target directory does not exist", func() {
				It("errs", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					_, err := dir.Find("snozzwangles", startDir)
					Ω(err).ShouldNot(BeNil())
				})
			})
		})
	})
})
