package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// Runs a command and transform it's output with customer parser.
// takes space seperated
func CommandOutput[T any](cmdArgs string, xform func(string) (T, error)) (*T, error) {
	parts := strings.Split(cmdArgs, " ")
	return CommandOutputA(parts, xform)
}

func CommandOutputA[T any](cmdArgs []string, xform func(string) (T, error)) (*T, error) {
	name := cmdArgs[0]
	args := cmdArgs[1:]
	stdout, stderr, retcode, err := runCommand(name, args)
	if err != nil {
		return nil, fmt.Errorf("%s: Error Running : %s", name, err.Error())
	}
	if retcode != 0 {
		return nil, fmt.Errorf("%s : Error running : ret code is not 0, is %d", name, retcode)
	}
	if stderr != "" {
		return nil, fmt.Errorf("%s : Error running : stderr :%s", name, stderr)
	}
	ret, err := xform(stdout)
	if err != nil {
		return nil, fmt.Errorf("%s : Error xforming : %s", name, err.Error())
	}
	return &ret, nil
}

// RunCommand executes a command and returns stdout, stderr, exit code, and error.
func runCommand(name string, args []string) (stdout, stderr string, retcode int, err error) {
	// Create command
	cmd := exec.Command(name, args...)

	// Capture stdout and stderr into buffers
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	// Run the command
	err = cmd.Run()

	// Default exit code = 0
	retcode = 0

	// If error occurred, extract exit code if possible
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				retcode = status.ExitStatus()
			}
		} else {
			// failed to start command (not just nonzero exit)
			retcode = -1
		}
	}

	// Return outputs as strings
	stdout = outBuf.String()
	stderr = errBuf.String()

	return
}
