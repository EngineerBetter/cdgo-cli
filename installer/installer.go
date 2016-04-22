package installer

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func Install(installTo string) error {
	if installTo == "" {
		return errors.New("Installation file was not specified")
	}
	return installBashFunctions(installTo)
}

func installBashFunctions(installTo string) error {
	installFileBytes, err := ioutil.ReadFile(installTo)
	if err != nil {
		return err
	}

	installFileContents := string(installFileBytes[:])
	functions := `
# https://github.com/EngineerBetter/cdgo
function cdgo {
  cd $(goto go -needle="$@")
}
function cdwork {
  cd $(goto work -needle="$@")
}
`
	if !strings.Contains(installFileContents, functions) {
		f, err := os.OpenFile(installTo, os.O_APPEND|os.O_WRONLY, 0666)
		defer f.Close()

		if err != nil {
			return err
		}

		_, err = f.WriteString(functions)
		if err != nil {
			return err
		}
	}

	return nil
}
