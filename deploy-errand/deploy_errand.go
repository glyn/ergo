package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/commandrunner/linux_command_runner"
	"github.com/glyn/ergo/cf"
)

func main() {
	fmt.Println("* Starting deploy errand")
	fmt.Printf("$SAMPLE=%s\n", os.Getenv("SAMPLE"))

	cf := cf.New(linux_command_runner.New())
	v, err := cf.DisplayCfVersion()
	if err != nil {
		fmt.Printf("cf.DisplayCfVersion failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(v)

	fmt.Println("* Finished deploy errand")
}
