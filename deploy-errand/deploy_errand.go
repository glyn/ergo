package main

import (
	"fmt"
	"github.com/glyn/ergo/cf"
	"code.cloudfoundry.org/commandrunner/linux_command_runner"
	"os"
)

func main() {
	fmt.Println("* Starting deploy errand")

	cf := cf.New(linux_command_runner.New())
	v, err := cf.DisplayCfVersion()
	if err != nil {
		fmt.Printf("cf.DisplayCfVersion failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(v)

	fmt.Println("* Finished deploy errand")
}
