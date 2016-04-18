package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/EngineerBetter/goto"
)

var _ = Describe("finder", func() {
	It("can be instantiated", func() {
		finder := new(Thing)
		Ω(finder).NotTo(BeNil())
	})

	Describe("find", func() {
		It("can be invoked", func() {
			finder := new(Thing)
			Ω(finder.Find("a", "b")).To(Equal("bar"))
		})
	})
})
