package cmd

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/go-errors/errors"
)

func CreateCmdConsole() CmdRunner {
	return &cmdConsole{}
}

type cmdConsole struct {
	cmdErr CmdError
}

func (c *cmdConsole) HasError() bool {
	return c.cmdErr.Err != nil
}

func (c *cmdConsole) SetError(cmdError CmdError) {
	c.cmdErr = cmdError
}

func (c *cmdConsole) LastError() CmdError {
	return c.cmdErr
}

func (c *cmdConsole) NewRunner() CmdRunner {
	return CreateCmdConsole()
}

func (c *cmdConsole) RunCmd(name string, args ...string) (string, CmdError) {
	stdout, stderr, err := RunCmd(name, args...)
	c.cmdErr = CmdError{Stderr: stderr, Err: err}
	return stdout, c.cmdErr
}

func (c *cmdConsole) RunInNewRunner(name string, args ...string) (string, CmdError) {
	newRunner := CreateCmdConsole()
	return newRunner.RunCmd(name, args...)
}

func RunCmd(name string, args ...string) (stdout string, stderr string, err error) {
	log.Printf("%v %v", name, args)
	cmd := exec.Command(name, args...)
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	cmdErr := cmd.Run()
	var e error = nil
	if cmdErr != nil {
		e = errors.Wrap(cmdErr, 1)
		log.Printf("Err output: %v %v", cmdErr, errBuffer.String())
	}
	return outBuffer.String(), errBuffer.String(), e
}
