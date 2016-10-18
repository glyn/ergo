package main

import (
	"github.com/glyn/ergo/cf"
	"github.com/glyn/ergo/util"
	"fmt"
	"os"
)

const defaultVcapSysDir = "/var/vcap/sys"
const defaultVcapRunDir = defaultVcapSysDir + "/run/" + jobName
const defaultVcapLogDir = defaultVcapSysDir + "/log" + jobName

const jobName = "deploy-service-broker"

func main() {
	fmt.Println("* Starting deploy errand")

	runDir, logDir := getVcapDirs()
	vcapUser, vcapGroup := getVcapUserInfo()

	setupVcapDirs([]string{runDir, logDir}, vcapUser, vcapGroup)
	displayCfVersion()
	cfTarget(os.Getenv("SYSTEM_DOMAIN"))
	cfAuth(os.Getenv("ADMIN_USER"), os.Getenv("ADMIN_PASSWORD"))

	fmt.Println("* Finished deploy errand")
}

func getVcapDirs() (string, string) {
	runDir := os.Getenv("VCAP_DIR_PREFIX") + defaultVcapRunDir
	logDir := os.Getenv("VCAP_DIR_PREFIX") + defaultVcapLogDir

	return runDir, logDir
}

func getVcapUserInfo() (string, string) {
	vcapUser := "vcap"

	if os.Getenv("VCAP_USER_NAME") != "" {
		vcapUser = os.Getenv("VCAP_USER_NAME")
	}

	vcapGroup := "vcap"

	if os.Getenv("VCAP_GROUP_NAME") != "" {
		vcapGroup = os.Getenv("VCAP_GROUP_NAME")
	}

	return vcapUser, vcapGroup
}

func setupVcapDirs(paths []string, userName string, groupName string) {
	err := util.CreateVcapDirs(paths, userName, groupName)

	if err != nil {
		os.Exit(1)
	}
}

func displayCfVersion() {
	output, err := cf.DisplayCfVersion()

	if err != nil {
		fmt.Printf("Failed to get CF version: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(output)
}

func cfTarget(systemDomain string) {
	output, err := cf.CfTarget(systemDomain)

	if err != nil {
		fmt.Printf("Failed to target CF endpoint: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(output)
}

func cfAuth(adminUserName string, adminPassword string) {
	output, err := cf.CfAuth(adminUserName, adminPassword)

	if err != nil {
		fmt.Printf("CF login failed: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(output)
}
