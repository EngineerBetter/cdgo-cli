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
	"strings"
)

var _ = Describe("goto", func() {
	var (
		cliPath string
		tmpDir  string
		envVars []string
	)

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/cdgo-cli")
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
		BeforeEach(func() {
			tmpDir, err := ioutil.TempDir("", "cdgo-cli-test")
			Ω(err).ShouldNot(HaveOccurred())
			envVars = os.Environ()

			for index, pair := range envVars {
				if strings.HasPrefix(pair, "GOPATH=") {
					envVars = append(envVars[:index], envVars[index+1:]...)
				}
			}

			envVars = append(envVars, "GOPATH="+tmpDir)

			os.MkdirAll(filepath.Join(tmpDir, "src", "github.com", "EngineerBetter", "Aardvark", "vendor", "github.com", "EngineerBetter", "Zebra"), 0700)
			os.MkdirAll(filepath.Join(tmpDir, "src", "github.com", "EngineerBetter", "Zebra"), 0700)
		})

		AfterEach(func() {
			Ω(os.RemoveAll(tmpDir)).Should(Succeed())
		})

		It("finds a dir not vendored anywhere", func() {
			command := exec.Command(cliPath, "-needle=Aardvark")
			command.Env = envVars
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))

			expectedOutput := filepath.Join(tmpDir, "src/github.com/EngineerBetter/Aardvark")
			Ω(session.Out).Should(Say(expectedOutput))
		})

		It("favours top-level dirs over those in vendor/", func() {
			command := exec.Command(cliPath, "-needle=Zebra")
			command.Env = envVars
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(0))

			expectedOutput := filepath.Join(tmpDir, "src/github.com/EngineerBetter/Zebra")
			Ω(session.Out).Should(Say(expectedOutput))
		})

		It("fails if the directory can't be found", func() {
			command := exec.Command(cliPath, "-needle=does-not-exist")
			command.Env = envVars
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
		})

		Describe("bash function installation", func() {
			It("adds the functions to the given file", func() {
				cliDir := filepath.Dir(cliPath)

				bashRcPath := filepath.Join(cliDir, ".bashrc")
				command := exec.Command(cliPath, "-install="+bashRcPath)
				command.Env = envVars
				session, err := Start(command, GinkgoWriter, GinkgoWriter)
				Ω(err).ShouldNot(HaveOccurred())
				Eventually(session).Should(Exit(0))
				Ω(session.Out).Should(Say("Added Bash functions to " + bashRcPath))

				command = exec.Command("bash", "-c", "export PATH=$PATH:. && cd "+cliDir+" && source .bashrc && cdgo cdgo-cli && pwd")
				session, err = Start(command, GinkgoWriter, GinkgoWriter)
				Ω(err).ShouldNot(HaveOccurred())
				Eventually(session).Should(Exit(0))
				expectedDir := filepath.Join(tmpDir, "src", "github.com", "EngineerBetter", "cdgo-cli")
				Ω(session.Out).Should(Say(expectedDir))
			})
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
})
