package cli

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)


func init(){
	rootCmd.AddCommand(restartCmd)
}


/*
func () {
	cmd := exec.Command("sudo", "systemctl", "restart", "proxyx")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to restart ProxyX:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("ProxyX restarted successfully")
}

*/

var restartCmd = &cobra.Command{
	Use: "restart",
	Short: "Reload ProxyX configuration",
	Run: func(cmd *cobra.Command, args []string) {
		restartProxyX()
	},
}



func restartProxyX() {
	fmt.Println("Restarting ProxyX service ....")

	var command *exec.Cmd

	switch os := runtime.GOOS; os {
	case "darwin" :
		command = exec.Command("sudo", "launchctl", "kickstart", "-k", "system/org.proxyx.service")
	case "linux" :
		command = exec.Command("sudo", "systemctl", "restart", "proxyx")
	default:
		fmt.Println("Unsupported OS:", os)
		return
	}

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to restart ProxyX:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("ProxyX restarted successfully")


}
