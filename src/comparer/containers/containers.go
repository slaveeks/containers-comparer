package containers

import (
	"awesomeProject/src/marks"
	"awesomeProject/src/utils"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type UtilType string

// Constants for running containers
const (
	Podman UtilType = "podman"
	Docker          = "docker"
)

type ContainerComparer struct {
	imageName string
	Utility   UtilType
}

func (c *ContainerComparer) RunContainer() {
	utility := string(c.Utility)

	utils.CallSystemctl("start", utility)

	t1 := time.Now()
	cmd := exec.Command(utility, "run", "-d", "-t", c.imageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while running image", err)
		return
	}
	containerID := strings.TrimSpace(string(out))
	t2 := time.Now()

	timer := marks.CreateTimeMark(t1, t2)

	ms := timer.TakeDiff()
	fmt.Printf("Container started in %d ms\n", ms)

	ramCmd := exec.Command(utility, "stats", "--no-stream", containerID, "--format", "{{.MemUsage}}")
	ramOut, err := ramCmd.Output()

	if err != nil {
		fmt.Println("Error while getting container utility stats", err)
		return
	}

	ramUsage := strings.TrimSpace(string(ramOut))

	num := utils.MegabytesToInt(ramUsage)

	fmt.Printf("Container RAM usage: %d\n", num)

	utils.CallSystemctl("stop", utility)
	utils.CallSystemctl("stop", utility+".socket")

}

func CreateContainersComparer(imageName string, utility UtilType) *ContainerComparer {
	return &ContainerComparer{imageName: imageName, Utility: utility}
}
