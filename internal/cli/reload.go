package cli

import (

	"github.com/spf13/cobra"
)

func (c *CLI) restartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "Reload ProxyX configuration",
		Run: func(cmd *cobra.Command, args []string) {
			c.Service.Restart()
		},
	}
}

/*

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Reload ProxyX configuration",
	Run: func(cmd *cobra.Command, args []string) {
		restartProxyX()
	},
}

func restartProxyX() {
	fmt.Println("Restarting ProxyX service ....")

	var command *exec.Cmd

	switch os := runtime.GOOS; os {
	case "darwin":
		command = exec.Command("sudo", "launchctl", "kickstart", "-k", "system/org.proxyx.service")
	case "linux":
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


*/