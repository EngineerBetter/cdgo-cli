package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var _ = Describe("goto", func() {
	var cliPath string
	var homeDir string
	var testDir string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/goto/workto")
		Ω(err).ShouldNot(HaveOccurred())

		usr, err := user.Current()
		Ω(err).ShouldNot(HaveOccurred())
		homeDir = usr.HomeDir
		testDir = filepath.Join(homeDir, "workspace", "goto-test-dir")
		err = os.MkdirAll(testDir, 0777)
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
		os.Remove(testDir)
	})

	It("requires an argument", func() {
		command := exec.Command(cliPath)
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit(1))
		Ω(session.Err).Should(Say("directory to look for was not specified"))
	})

	It("finds a test directory", func() {
		command := exec.Command(cliPath, "goto-test-dir")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session, "5s").Should(Exit())

		expectedOutput := testDir
		Ω(session.Out).Should(Say(expectedOutput))
	})
})
