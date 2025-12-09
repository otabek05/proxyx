package cli

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "Stops running proxyx service",
	Run: func(cmd *cobra.Command, args []string) {
		stopProxyX()
	},
}

func stopProxyX() {
	fmt.Println("Stopping ProxyX service...")
	var command *exec.Cmd
	switch os := runtime.GOOS; os {
	case "darwin":
		command = exec.Command("sudo", "launchctl", "stop", "org.proxyx.service")
	case "linux":
		command = exec.Command("sudo", "systemctl", "stop", "proxyx")
	default:
		fmt.Println("Unsupported OS:", os)
		return
	}

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to stop ProxyX:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("ProxyX stopped successfully")
}