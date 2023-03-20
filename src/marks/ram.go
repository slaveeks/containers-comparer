package marks

import (
	"fmt"
	"os/exec"
)

func GetCurrentMemoryUsage() {
	cmd := exec.Command("free")

	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}
