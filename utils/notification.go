package utils

import (
	"os/exec"
)

func SendDesktopNotification(summary, body string) error {
	// Use the notify-send command to send a desktop notification
	cmd := exec.Command("notify-send", summary, body)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
