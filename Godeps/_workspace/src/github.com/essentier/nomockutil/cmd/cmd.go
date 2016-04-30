package cmd

type CmdError struct {
	Stderr string
	Err    error
}

func (ce CmdError) Error() string {
	return ce.Err.Error() + " stderr: " + ce.Stderr
}

type CmdRunner interface {
	SetError(cmdError CmdError)
	RunCmd(name string, args ...string) (string, CmdError)
	RunInNewRunner(name string, args ...string) (string, CmdError)
	NewRunner() CmdRunner
	LastError() CmdError
	HasError() bool
}
