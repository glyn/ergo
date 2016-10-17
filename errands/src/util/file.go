package util

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func CreateVcapDirs(paths []string, userName string, groupName string) error {
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
