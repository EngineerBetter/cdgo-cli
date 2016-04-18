package finder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/EngineerBetter/goTo/finder"
)

var _ = Describe("finder", func() {
	It("can be instantiated", func() {
		finder := new(finder.Thing)
		Ω(finder).NotTo(BeNil())
	})
})
