package utils

import (
	"os/exec"
	"strings"
)

func CheckStatus(Service string) (bool, error) {
	cmd := exec.Command("sudo", "systemctl", "is-active", Service)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	} else {
		result := strings.TrimSpace(string(output))
		if result == "active" {
			return true, nil
		} else {
			return false, nil
		}

	}
}
