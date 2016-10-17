package cf

import "os/exec"

func DisplayCfVersion() (string, error) {
	cmd := exec.Command("cf", "-v")
	output, err := cmd.Output()

	return string(output), err
}

func CfTarget(systemDomain string) (string, error) {
	apiEndpoint := "api." + systemDomain

	// TODO: make ssl validation conditional
	cmd := exec.Command("cf", "api", apiEndpoint, "--skip-ssl-validation")
	output, err := cmd.Output()

	return string(output), err
}

func CfAuth(userName string, password string) (string, error) {
	cmd := exec.Command("cf", "auth", userName, password)
	output, err := cmd.Output()

	return string(output), err
}
