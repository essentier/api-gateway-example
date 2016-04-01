package cmd

func CreateFailFastCmdRunner(innerCmdRunner CmdRunner) CmdRunner {
	return &failFastCmdRunner{CmdRunner: innerCmdRunner}
}

type failFastCmdRunner struct {
	CmdRunner
}

func (c *failFastCmdRunner) NewRunner() CmdRunner {
	return CreateFailFastCmdRunner(c.CmdRunner.NewRunner())
}

func (c *failFastCmdRunner) RunInNewRunner(name string, args ...string) (string, CmdError) {
	newRunner := c.NewRunner()
	return newRunner.RunCmd(name, args...)
}

func (c *failFastCmdRunner) RunCmd(name string, args ...string) (string, CmdError) {
	if c.HasError() {
		return "", c.LastError()
	}

	return c.CmdRunner.RunCmd(name, args...)
}

// func (c *failFastCmdRunner) SetError(cmdError CmdError) {
// 	c.cmdErr = cmdError
// }

// func (c *failFastCmdRunner) LastError() CmdError {
// 	return c.cmdErr
// }
