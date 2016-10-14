package helper

import (
	"os/exec"

	"code.cloudfoundry.org/commandrunner"
)

func Echo(cmdRunner commandrunner.CommandRunner) {
	cmd := exec.Command("echo", "shelled out")
	cmdRunner.Run(cmd)
}
