package main

import (
	"cf"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

const defaultVcapSysDir = "/var/vcap/sys"
const defaultVcapRunDir = defaultVcapSysDir + "/run/" + jobName
const defaultVcapLogDir = defaultVcapSysDir + "/log" + jobName

const jobName = "deploy-service-broker"

func main() {
	fmt.Println("* Starting deploy errand")

	runDir := os.Getenv("VCAP_DIR_PREFIX") + defaultVcapRunDir
	logDir := os.Getenv("VCAP_DIR_PREFIX") + defaultVcapLogDir

	vcapUser := "vcap"

	if os.Getenv("VCAP_USER_NAME") != "" {
		vcapUser = os.Getenv("VCAP_USER_NAME")
	}

	vcapGroup := "vcap"

	if os.Getenv("VCAP_GROUP_NAME") != "" {
		vcapGroup = os.Getenv("VCAP_GROUP_NAME")
	}

	setupVcapDirs([]string{runDir, logDir}, vcapUser, vcapGroup)
	displayCfVersion()
	cfTarget(os.Getenv("SYSTEM_DOMAIN"))
	cfAuth(os.Getenv("ADMIN_USER"), os.Getenv("ADMIN_PASSWORD"))

	fmt.Println("* Finished deploy errand")
}

func setupVcapDirs(paths []string, userName string, groupName string) {
	err := createVcapDirs(paths, userName, groupName)

	if err != nil {
		os.Exit(1)
	}
}

func createVcapDirs(paths []string, userName string, groupName string) error {
	err := mkdir(paths)

	if err != nil {
		return err
	}

	err = chown(paths, userName, groupName)

	if err != nil {
		return err
	}

	return nil
}

func mkdir(paths []string) error {
	for _, path := range paths {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Make directory failed: %s\n", err)
			return err
		}
	}

	return nil
}

func chown(paths []string, userName string, groupName string) error {
	userId, err := getUserId(userName)

	if err != nil {
		return err
	}

	groupId, err := getGroupId(groupName)

	if err != nil {
		return err
	}

	for _, path := range paths {
		err := os.Chown(path, userId, groupId)

		if err != nil {
			fmt.Printf("Chown failed: %s\n", err)
		}
	}

	return nil
}

func getUserId(userName string) (int, error) {
	u, err := user.Lookup(userName)

	if err != nil {
		fmt.Printf("Failed to obtain UID: %s\n", err)
		return -1, err
	}

	userId, _ := strconv.Atoi(u.Uid)

	return userId, nil
}

func getGroupId(groupName string) (int, error) {
	g, err := user.LookupGroup(groupName)

	if err != nil {
		fmt.Printf("Failed to obtain GID: %s\n", err)
		return -1, err
	}

	groupId, _ := strconv.Atoi(g.Gid)

	return groupId, nil
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
