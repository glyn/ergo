package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/commandrunner/linux_command_runner"
	"github.com/glyn/ergo/cf"
	"io/ioutil"
	"encoding/json"
)

type DeploymentProperties struct {
	SpringCloudBroker SpringCloudBrokerProperties `json:"spring_cloud_broker"`
}

type SpringCloudBrokerProperties struct {
	CF                    CFProperties
}

type CFProperties struct {
	Sample string
}

func main() {
	fmt.Println("* Starting deploy errand")

	jsonPath := os.Args[1]
	j, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Failed to read JSON properties file: %s\n", err)
		os.Exit(1)
	}

	var props DeploymentProperties
	err = json.Unmarshal(j, &props)
	if err != nil {
		fmt.Printf("Failed to unmarshall JSON properties: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("spring_cloud_broker.cf.sample=%s\n", props.SpringCloudBroker.CF.Sample)

	cf := cf.New(linux_command_runner.New())
	v, err := cf.DisplayCfVersion()
	if err != nil {
		fmt.Printf("cf.DisplayCfVersion failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(v)

	fmt.Println("* Finished deploy errand")
}
