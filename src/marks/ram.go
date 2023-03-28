package marks

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetCurrentMemoryUsage(cmd *exec.Cmd) {
	pid := fmt.Sprintf("%d", cmd.Process.Pid)
	output, err := exec.Command("ps", "-p", pid, "-o", "rss=").Output()
	if err != nil {
		fmt.Println("Error getting RAM usage: ", err)
		os.Exit(1)
	}

	// Convert the output to an integer and print it
	ramUsage, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		fmt.Println("Error converting RAM usage to integer: ", err)
		os.Exit(1)
	}
	fmt.Println("RAM usage during build: ", ramUsage, "KB")
}
