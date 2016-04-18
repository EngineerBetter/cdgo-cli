package main_test

import (
	. "github.com/EngineerBetter/goto"
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
		Ω(finder).NotTo(BeNil())
		cwd, _ = os.Getwd()
	})

	Describe("find", func() {
		Context("when the starting directory does not exist", func() {
			It("returns an error", func() {
				startDir := filepath.Join(cwd, "not-exist")
				_, err := finder.Find("needle", startDir)
				Ω(err).NotTo(BeNil(), "Error should be returned for non-existant directory")
			})
		})

		Context("when the starting directory exists", func() {
			Context("and the target directory exists", func() {
				It("returns the absolute path to the directory", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					result, err := finder.Find("root", startDir)
					Ω(err).To(BeNil())
					Ω(result).To(Equal(cwd + "/test-fixtures/root"))
				})
			})

			Context("but the target directory does not exist", func() {
				It("errs", func() {
					startDir := filepath.Join(cwd, "test-fixtures")
					_, err := finder.Find("snozzwangles", startDir)
					Ω(err).ToNot(BeNil())
				})
			})
		})
	})
})
