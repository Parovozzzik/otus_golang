package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	for fileName, envValue := range env {
		os.Setenv(fileName, envValue.Value)
	}

	command.Run()

	return command.ProcessState.ExitCode()
}
