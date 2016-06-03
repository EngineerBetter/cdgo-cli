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
	"strings"
)

var _ = Describe("goto", func() {
	var (
		cliPath string
		gopath  string
	)

	BeforeSuite(func() {
		var err error
		gopath = os.Getenv("GOPATH")
		currentLocation, err := filepath.Abs(os.Args[0])
		Ω(err).ShouldNot(HaveOccurred())
		splitter := strings.Split(currentLocation, "github.com")
		splitter = strings.Split(splitter[1], "/_test")
		buildPath := filepath.Join("github.com", splitter[0])
		cliPath, err = Build(buildPath)
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
		Context("when there is a project in the root of GOPATH that contains the target dir within vendor", func() {
			BeforeEach(func() {
				newDirPath := filepath.Join(gopath, "src/cdgo-example/vendor/github.com/EngineerBetter/cdgo-example2")
				err := os.MkdirAll(newDirPath, 0777)
				Ω(err).ShouldNot(HaveOccurred())
				newDirPath = filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo-example2")
				err = os.MkdirAll(newDirPath, 0777)
				Ω(err).ShouldNot(HaveOccurred())
			})

			AfterEach(func() {
				newDirPath := filepath.Join(gopath, "src/cdgo-example")
				err := os.RemoveAll(newDirPath)
				Ω(err).ShouldNot(HaveOccurred())
				newDirPath = filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo-example2")
				err = os.RemoveAll(newDirPath)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("finds this directory", func() {
				Ω(gopath).ShouldNot(BeZero())

				command := exec.Command(cliPath, "-needle=cdgo-example2")
				session, err := Start(command, GinkgoWriter, GinkgoWriter)
				Ω(err).ShouldNot(HaveOccurred())
				Eventually(session).Should(Exit(0))

				expectedOutput := filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo-example2")
				Ω(session.Out).Should(Say(expectedOutput))
			})
		})

		Context("when vendor is of no concern", func() {
			It("finds this directory", func() {
				Ω(gopath).ShouldNot(BeZero())

				command := exec.Command(cliPath, "-needle=cdgo-cli")
				session, err := Start(command, GinkgoWriter, GinkgoWriter)
				Ω(err).ShouldNot(HaveOccurred())
				Eventually(session).Should(Exit(0))

				expectedOutput := filepath.Join(gopath, "src/github.com/EngineerBetter/cdgo-cli")
				Ω(session.Out).Should(Say(expectedOutput))
			})
		})

		It("fails if the directory can't be found", func() {
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
			cliPath, err = Build("github.com/EngineerBetter/cdgo-cli")
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
			Eventually(session).Should(Exit(1))
		})
	})

	Describe("bash function installation", func() {
		It("adds the functions to the given file", func() {
			cliDir := filepath.Dir(cliPath)

			bashRcPath := filepath.Join(cliDir, ".bashrc")
			command := exec.Command(cliPath, "-install="+bashRcPath)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))
			Ω(session.Out).Should(Say("Added Bash functions to " + bashRcPath))

			command = exec.Command("bash", "-c", "export PATH=$PATH:. && cd "+cliDir+" && source .bashrc && cdgo cdgo-cli && pwd")
			session, err = Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))
			expectedDir := filepath.Join(gopath, "src", "github.com", "EngineerBetter", "cdgo-cli")
			Ω(session.Out).Should(Say(expectedDir))
		})
	})
})
