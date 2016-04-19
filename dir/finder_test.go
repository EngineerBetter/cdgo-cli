package dir_test

import (
	. "github.com/EngineerBetter/goto/dir"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"path/filepath"
)

var _ = Describe("finder", func() {
	var (
		cwd    string
		finder DirectoryFinder
	)

	BeforeEach(func() {
		finder = new(RecursiveFinder)
		Ω(finder).ShouldNot(BeNil())
		cwd, _ = os.Getwd()
	})

	Describe("find", func() {
		Context("when the starting directory does not exist", func() {
			It("returns an error", func() {
				startDir := filepath.Join(cwd, "not-exist")
				_, err := finder.Find("needle", startDir)
				Ω(err).ShouldNot(BeNil(), "Error should be returned for non-existant directory")
			})
		})

		Context("when the starting directory exists", func() {
			Context("and the target directory exists", func() {
				It("returns the absolute path to the highest matching directory", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					result, err := finder.Find("root", startDir)
					Ω(err).Should(BeNil())
					Ω(result).Should(Equal(cwd + "/test-fixtures/root"))
				})
			})

			Context("but the target directory does not exist", func() {
				It("errs", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					_, err := finder.Find("snozzwangles", startDir)
					Ω(err).ShouldNot(BeNil())
				})
			})
		})
	})
})
