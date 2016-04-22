package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
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
		Ω(session.Err).Should(Say("-needle must be provided"))
	})

	Describe("switching to Go dirs", func() {
		It("finds this directory", func() {
			gopath := os.Getenv("GOPATH")
			Ω(gopath).ShouldNot(BeZero())

			command := exec.Command(cliPath, "-needle=cdgo")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))

			expectedOutput := filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo")
			Ω(session.Out).Should(Say(expectedOutput))
		})

		It("fails if the directory can't be found", func() {
			gopath := os.Getenv("GOPATH")
			Ω(gopath).ShouldNot(BeZero())

			command := exec.Command(cliPath, "-needle=does-not-exist")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
		})
	})

	Describe("switching to Work dirs", func() {
		var homeDir string
		var testDir string

		BeforeEach(func() {
			var err error
			cliPath, err = Build("github.com/EngineerBetter/cdgo/goto")
			Ω(err).ShouldNot(HaveOccurred())

			usr, err := user.Current()
			Ω(err).ShouldNot(HaveOccurred())
			homeDir = usr.HomeDir
			testDir = filepath.Join(homeDir, "workspace", "goto-test-dir")
			err = os.MkdirAll(testDir, 0777)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("finds a test directory", func() {
			command := exec.Command(cliPath, "-haystackType=work", "-needle=goto-test-dir")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))

			expectedOutput := testDir
			Ω(session.Out).Should(Say(expectedOutput))
		})

		It("fails if the directory can't be found", func() {
			command := exec.Command(cliPath, "-haystackType=work", "-needle=does-not-exist")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session, "30s").Should(Exit(1))
		})
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
			Eventually(session).Should(Exit(0))
			Ω(session.Out).Should(Say("Added Bash functions to " + tempFilePath))

			bytesAfter, err := ioutil.ReadFile(tempFilePath)
			Ω(err).ShouldNot(HaveOccurred())

			functions := `
# https://github.com/EngineerBetter/cdgo
function cdgo {
  cd $(goto go "$@")
}
function cdwork {
  cd $(goto work "$@")
}
`
			stringAfter := string(bytesAfter[:])
			Ω(stringAfter).Should(ContainSubstring(functions))
		})
	})
})
