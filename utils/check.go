package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckStatus(Service string) (bool, error) {
	cmd := exec.Command("sudo", "systemctl", "is-active", Service)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Get the exit status code
			statusCode := exitError.ExitCode()
			// Need to handle the Exit codes in a better way
			if statusCode == 3 {
				return false, nil
			} else {
				fmt.Printf("Error: Service %s exited with status code %d\n", Service, statusCode)
			}
			return false, err
		}
	} else {
		result := strings.TrimSpace(string(output))
		if result == "active" {
			return true, nil
		}
	}
	return false, nil
}
