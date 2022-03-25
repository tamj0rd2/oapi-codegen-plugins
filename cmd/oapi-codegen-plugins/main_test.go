package main_test

import (
	"context"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("it prints the given input", func(t *testing.T) {
		const inputCode = `package testing

type GetTerminalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AuthorizationGroupCreated bool
		BusinessCreated           bool
		BusinessId *string
	}
}
`

		pluginCommand := exec.CommandContext(context.Background(), "go", "run", "main.go")
		pluginCommand.Stdin = strings.NewReader(inputCode)

		cliOutputBytes, err := pluginCommand.Output()
		if err != nil {
			t.Fatal(err)
		}

		cliOutput := string(cliOutputBytes)

		if strings.Compare(cliOutput, inputCode) != 0 {
			t.Errorf("Expected output to be %s, got %s", inputCode, cliOutput)
		}

		t.Log("Tests worked fine")
		t.Logf(string(cliOutputBytes))
	})
}
