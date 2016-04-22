package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var _ = Describe("goto", func() {
	var cliPath string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/cdgo/goto")
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
		Ω(session.Err).Should(Say("directory to look for was not specified"))
	})

	It("finds this directory", func() {
		gopath := os.Getenv("GOPATH")
		Ω(gopath).ShouldNot(BeZero())

		command := exec.Command(cliPath, "cdgo")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(Exit())

		expectedOutput := filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo")

		Ω(session.Out).Should(Say(expectedOutput))
	})

	Describe("bash function installation", func() {
		It("adds the functions to the given file", func() {
			tempFile, err := ioutil.TempFile("", ".bashrc")
			Ω(err).ShouldNot(HaveOccurred())
			tempFilePath := tempFile.Name()
			tempFile.Close()
			defer os.Remove(tempFilePath)

			command := exec.Command(cliPath, "-install="+tempFilePath)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit())
			Ω(session.Out).Should(Say("Added Bash functions to " + tempFilePath))

			bytesAfter, err := ioutil.ReadFile(tempFilePath)
			Ω(err).ShouldNot(HaveOccurred())

			functions := `
# https://github.com/EngineerBetter/cdgo
function cdgo {
  cd $(goto "$@")
}
function cdwork {
  cd $(workto "$@")
}
`

			stringAfter := string(bytesAfter[:])
			Ω(stringAfter).Should(ContainSubstring(functions))
		})
	})
})
