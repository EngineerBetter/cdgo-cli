package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/EngineerBetter/goTo"
)

var _ = Describe("finder", func() {
	It("can be instantiated", func() {
		finder := new(Thing)
		Ω(finder).NotTo(BeNil())
	})
})
