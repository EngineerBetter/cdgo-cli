package installer_test

import (
	"github.com/EngineerBetter/cdgo/goto/installer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Describe("installer", func() {
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
function cdgo {
  cd $(goto "$@")
}
function cdwork {
  cd $(workto "$@")
}
`

		err = installer.Install(tempFilePath)
		Ω(err).ShouldNot(HaveOccurred())

		bytesAfter, err := ioutil.ReadFile(tempFilePath)
		Ω(err).ShouldNot(HaveOccurred())

		stringAfter := string(bytesAfter[:])
		Ω(stringAfter).Should(ContainSubstring(functions))
		Ω(stringAfter).Should(ContainSubstring(stringBefore))
	})
})
