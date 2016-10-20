package cf

import (
	"os/exec"
	"code.cloudfoundry.org/commandrunner"
	"bytes"
)

type CF interface {
	DisplayCfVersion() (string, error)
}

type cf struct {
	commandRunner commandrunner.CommandRunner
}

func New(commandRunner commandrunner.CommandRunner) CF {
	return &cf{commandRunner: commandRunner}
}

func (cf *cf) DisplayCfVersion() (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("cf", "-v")
	cmd.Stdout = &stdout

	err := cf.commandRunner.Run(cmd)
	if (err != nil) {
		return "", err
	}

	return stdout.String(), err
}
