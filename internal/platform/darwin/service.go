package darwin

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {

	return nil
}

func (s *Service) Restart() error {
	command := exec.Command("sudo", "launchctl", "stop", "org.proxyx.service")
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to stop ProxyX:", err)
		fmt.Println(string(output))
		return err
	}

	return nil
}

func (s *Service) Status() error {
	cmdCheck := exec.Command("sudo", "launchctl", "print", "system/org.proxyx.service")
	output, err := cmdCheck.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ProxyX is not running")
	}

	status := string(output)
	lines := strings.Split(status, "\n")
	var pid int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "pid =") {
			fmt.Sscanf(line, "pid = %d", &pid)
			break
		}
	}

	if pid == 0 {
		return fmt.Errorf("ProxyX is not running")
	}

	psCmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "pid,pcpu,pmem,etime")
	psOutput, _ := psCmd.CombinedOutput()
	psLines := strings.Split(strings.TrimSpace(string(psOutput)), "\n")
	if len(psLines) < 2 {
		return fmt.Errorf("Failed to get process stats")
	}

	psFields := strings.Fields(psLines[1])
	if len(psFields) < 4 {
		return fmt.Errorf("Failed to get process stats")
	}

	fmt.Println("ProxyX is running (launchd service)")
	fmt.Println("PID       CPU%    MEM%    Uptime")
	fmt.Printf("%-9s %-7s %-7s %-7s\n", psFields[0], psFields[1], psFields[2], psFields[3])
	return nil
}
