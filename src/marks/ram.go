package marks

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetCurrentMemoryUsage() {
	cmd := exec.Command("free")

	output, _ := cmd.CombinedOutput()

	for n, line := range strings.Split(strings.TrimSuffix(string(output), "\n"), "\n") {
		if n == 1 {
			a := strings.Fields(line)
			fmt.Println(a[2])
		}
	}
}
