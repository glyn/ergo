package cf

 import ( "fmt"
	"os/exec"
	"os"
)

func DisplayCfVersion() {
	cmd := exec.Command("cf", "-v")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Failed to get CF version: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func CfTarget(systemDomain string) {
	apiEndpoint := "api." + systemDomain

	cmd := exec.Command("cf", "api", apiEndpoint, "--skip-ssl-validation")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Failed to target CF endpoint: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func CfAuth(userName string, password string) {
	cmd := exec.Command("cf", "auth", userName, password)
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("CF login failed: %s\n", output)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

