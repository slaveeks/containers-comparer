package utils

import (
	"fmt"
	"os/exec"
)

func CallSystemctl(command string, arg string) {
	var cmd *exec.Cmd

	cmd = exec.Command("systemctl", command, arg)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error while executing systemctl command " + command + " " + arg)
	}
}
