package cli

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if ProxyX is running or not",
	Run: func(cmd *cobra.Command, args []string) {
		checkStatus()
	},
}

func checkStatus() {
	switch runtime.GOOS {
	case "linux":
		checkStatusLinux()
	case "darwin":
		checkStatusMacOS()
	default:
		fmt.Println("Unsupported OS")
	}
}

func checkStatusLinux() {
	cmdCheck := exec.Command("systemctl", "is-active", "proxyx")
	output, _ := cmdCheck.CombinedOutput()
	status := strings.TrimSpace(string(output))

	if status != "active" {
		fmt.Println("ProxyX is not running")
		return
	}

	fmt.Println("ProxyX is running (systemd service)")

	cmdPID := exec.Command("systemctl", "show", "proxyx", "-p", "MainPID")
	pidOutput, _ := cmdPID.CombinedOutput()
	var pid int
	fmt.Sscanf(string(pidOutput), "MainPID=%d", &pid)
	if pid == 0 {
		fmt.Println("Cannot find ProxyX PID")
		return
	}

	psCmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "pid,pcpu,pmem,etime=", "--no-headers")
	psOutput, _ := psCmd.CombinedOutput()
	psFields := strings.Fields(string(psOutput))
	if len(psFields) < 4 {
		fmt.Println("Failed to get process stats")
		return
	}

	fmt.Println("PID       CPU%    MEM%    Uptime")
	fmt.Printf("%-9s %-7s %-7s %-7s\n", psFields[0], psFields[1], psFields[2], psFields[3])
}

func checkStatusMacOS() {
	cmdCheck := exec.Command("launchctl", "list", "org.proxyx.service")
	output, _ := cmdCheck.CombinedOutput()
	status := strings.TrimSpace(string(output))

	if status == "" {
		fmt.Println("ProxyX is not running")
		return
	}

	lines := strings.Split(status, "\n")
	if len(lines) == 0 {
		fmt.Println("ProxyX is not running")
		return
	}

	fields := strings.Fields(lines[0])
	if len(fields) < 3 {
		fmt.Println("Failed to parse launchctl output")
		return
	}

	pidStr := fields[0]
	pid, err := strconv.Atoi(pidStr)
	if err != nil || pid == 0 {
		fmt.Println("ProxyX is not running")
		return
	}

	psCmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "pid,pcpu,pmem,etime=", "--no-headers")
	psOutput, _ := psCmd.CombinedOutput()
	psFields := strings.Fields(string(psOutput))
	if len(psFields) < 4 {
		fmt.Println("Failed to get process stats")
		return
	}

	fmt.Println("ProxyX is running (launchd service)")
	fmt.Println("PID       CPU%    MEM%    Uptime")
	fmt.Printf("%-9s %-7s %-7s %-7s\n", psFields[0], psFields[1], psFields[2], psFields[3])
}
