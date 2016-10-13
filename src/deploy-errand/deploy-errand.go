package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"cf"
)

const packageName = "spring-cloud-service-broker"
const packagePath = "/var/vcap/packages/" + packageName
const jobName = "deploy-service-broker"
const runDir = "/var/vcap/sys/run/" + jobName
const logDir = "/var/vcap/sys/log/" + jobName
const vcapUser = "vcap"
const vcapGroup = "vcap"

var adminUserName = os.Getenv("ADMIN_USER")
var adminPassword = os.Getenv("ADMIN_PASSWORD")
var systemDomain = os.Getenv("SYSTEM_DOMAIN")

func main() {
	fmt.Println("Starting deploy errand")

	setupVcapDirs()

	cf.DisplayCfVersion()
	cf.CfTarget(systemDomain)
	cf.CfAuth(adminUserName, adminPassword)
}

func setupVcapDirs() {
	mkdir([]string{runDir, logDir})
	chown([]string{runDir, logDir}, vcapUser, vcapGroup)
}

func mkdir(paths []string) {
	for _, path := range paths {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Make directory failed: %s\n", err)
			os.Exit(1)
		}
	}
}

func chown(paths []string, userName string, groupName string) {
	userId := getUserId(userName)
	groupId := getGroupId(groupName)

	for _, path := range paths {
		err := os.Chown(path, userId, groupId)

		if err != nil {
			fmt.Printf("Chown failed: %s\n", err)
			os.Exit(1)
		}
	}
}

func getUserId(userName string) int {
	u, err := user.Lookup(userName)

	if err != nil {
		fmt.Printf("Failed to obtain UID for user name: %s\n", userName, err)
		os.Exit(1)
	}

	userId, _ := strconv.Atoi(u.Uid)

	return userId
}

func getGroupId(groupName string) int {
	g, err := user.LookupGroup(groupName)

	if err != nil {
		fmt.Printf("Failed to obtain GID for group name: %s\n", groupName, err)
		os.Exit(1)
	}

	groupId, _ := strconv.Atoi(g.Gid)

	return groupId
}
