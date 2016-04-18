package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/EngineerBetter/goto"
)

var _ = Describe("finder", func() {
	var finder DirectoryFinder

	BeforeEach(func() {
		finder = new(RecursiveFinder)
		Ω(finder).NotTo(BeNil())
	})

	Describe("find", func() {
		It("can be invoked", func() {
			finder := new(RecursiveFinder)
			Ω(finder.Find("a", "b")).To(Equal("bar"))
		})
	})
})
