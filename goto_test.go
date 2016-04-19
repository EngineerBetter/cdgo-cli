package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"os/exec"
)

var _ = Describe("goto", func() {
	var cliPath string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/goto")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	It("requires an argument", func() {
		command := exec.Command(cliPath)
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Out).Should(Say("directory to look for was not specified"))
	})
})
