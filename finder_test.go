package main_test

import (
	. "github.com/EngineerBetter/goto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"path/filepath"
)

var _ = Describe("finder", func() {
	var finder DirectoryFinder

	BeforeEach(func() {
		finder = new(RecursiveFinder)
		Ω(finder).NotTo(BeNil())
	})

	Describe("find", func() {
		Context("when the starting directory does not exist", func() {
			It("returns an error", func() {
				cwd, err := os.Getwd()
				startDir := filepath.Join(cwd, "not-exist")
				_, err = finder.Find("needle", startDir)
				Ω(err).NotTo(BeNil(), "should err")
			})
		})
	})
})
