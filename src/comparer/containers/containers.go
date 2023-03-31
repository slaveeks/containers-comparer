package containers

import (
	"awesomeProject/src/marks"
	"awesomeProject/src/utils"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
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

	re := regexp.MustCompile(`^\d+`)

	match := re.FindString(ramUsage)

	num, _ := strconv.Atoi(match)

	fmt.Printf("Container RAM usage: %d\n", num)

	utils.CallSystemctl("stop", utility)
	utils.CallSystemctl("stop", utility+".socket")

}
