package cf

import (
	"bytes"
	"code.cloudfoundry.org/commandrunner"
	"os/exec"
	"code.cloudfoundry.org/commandrunner/linux_command_runner"
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

func DefaultNew() CF {
	return &cf{commandRunner: linux_command_runner.New()}
}

func (cf *cf) DisplayCfVersion() (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("cf", "-v")
	cmd.Stdout = &stdout

	err := cf.commandRunner.Run(cmd)
	if err != nil {
		return "", err
	}

	return stdout.String(), err
}
