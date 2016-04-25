package bash

import (
	"errors"
	"fmt"
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
	if _, err := os.Stat(installTo); os.IsNotExist(err) {
		_, err = os.Create(installTo)

		if err != nil {
			return errors.New("Could not create new file " + installTo)
		}
		fmt.Println("Created file " + installTo)
	}

	installFileBytes, err := ioutil.ReadFile(installTo)
	if err != nil {
		return err
	}

	installFileContents := string(installFileBytes[:])
	functions := `
# https://github.com/EngineerBetter/cdgo-cli
function cdgo { cd $(cdgo-cli -needle="$@") ; }
function cdwork { cd $(cdgo-cli -haystackType=work -needle="$@") ; }
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
