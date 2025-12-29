package cli

import (
	"github.com/spf13/cobra"
)

func (c *CLI) stopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops running proxyx service",
		RunE: func(cmd *cobra.Command, args []string) error  {
			return c.Service.Stop()
		},
	}

}


/*
var stopCmd = &cobra.Command{
	Use:   "stop",
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
*/