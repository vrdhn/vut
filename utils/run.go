package utils

import (
	"bytes"
	"os/exec"
	"syscall"
)

// RunCommand executes a command and returns stdout, stderr, exit code, and error.
func RunCommand(name string, args []string) (stdout, stderr string, retcode int, err error) {
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

func CommandOutput(name string, args ...string) (string, bool) {

	stdout, stderr, retcode, err := RunCommand(name, args)
	if err != nil {
		LogError(name + ": Error running : " + err.Error())
		return "", false
	}
	if retcode != 0 {
		LogError(name + ": Error running : ret code is not 0, is " + string(retcode))
		return "", false
	}
	if stderr != "" {
		LogError(name + ": Error running : stderr :" + stderr)
		return "", false
	}
	return stdout, true
}
