package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

func CallSystemctl(command string, arg string) {
	var cmd *exec.Cmd

	cmd = exec.Command("systemctl", command, arg)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error while executing systemctl command " + command + " " + arg)
	}
}

func MegabytesToInt(str string) int {
	re := regexp.MustCompile(`^\d+`)

	match := re.FindString(str)

	num, _ := strconv.Atoi(match)

	return num
}
