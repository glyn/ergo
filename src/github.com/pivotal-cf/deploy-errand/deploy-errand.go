package main

import (
	"fmt"
	"github.com/pivotal-cf/deploy-errand/helper"

	"code.cloudfoundry.org/commandrunner/linux_command_runner"
)

func main() {
	fmt.Println("hello from golang")
	helper.Echo(linux_command_runner.New())
}
