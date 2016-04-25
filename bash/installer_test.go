package bash_test

import (
	"github.com/EngineerBetter/cdgo/bash"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Describe("bash", func() {
	Context("when the file exists", func() {
		It("adds Bash functions to the specified file", func() {
			workDir, err := os.Getwd()
			Ω(err).ShouldNot(HaveOccurred())
			testFilePath := filepath.Join(workDir, "test-fixtures", ".bashrc")
			bytesBefore, err := ioutil.ReadFile(testFilePath)
			Ω(err).ShouldNot(HaveOccurred())

			tempFile, err := ioutil.TempFile("", ".bashrc")
			Ω(err).ShouldNot(HaveOccurred())

			_, err = tempFile.Write(bytesBefore)
			Ω(err).ShouldNot(HaveOccurred())
			tempFilePath := tempFile.Name()
			tempFile.Close()
			defer os.Remove(tempFilePath)

			stringBefore := string(bytesBefore[:])

			functions := `
# https://github.com/EngineerBetter/cdgo
function cdgo { cd $(goto -needle="$@") ; }
function cdwork { cd $(goto -haystackType=work -needle="$@") ; }
`

			err = bash.Install(tempFilePath)
			Ω(err).ShouldNot(HaveOccurred())

			bytesAfter, err := ioutil.ReadFile(tempFilePath)
			Ω(err).ShouldNot(HaveOccurred())

			stringAfter := string(bytesAfter[:])
			Ω(stringAfter).Should(ContainSubstring(functions))
			Ω(stringAfter).Should(ContainSubstring(stringBefore))
		})
	})

	Context("when the file does not exist", func() {
		It("creates the file with the Bash functions in", func() {
			dir, err := ioutil.TempDir("", "prefix")
			Ω(err).ShouldNot(HaveOccurred())

			file := filepath.Join(dir, ".bashrc")
			_, err = os.Stat(file)
			Ω(err).Should(HaveOccurred())
			err = bash.Install(file)
			Ω(err).ShouldNot(HaveOccurred())
			_, err = os.Stat(file)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
