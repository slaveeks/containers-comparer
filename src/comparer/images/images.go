package images

import (
	"awesomeProject/src/marks"
	"awesomeProject/src/utils"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type UtilType string

type CompareData struct {
	buildTime      int
	buildRamUsage  int
	buildImageSize int
}

// Constants for building images
const (
	Podman  UtilType = "podman"
	Docker           = "docker"
	Img              = "img"
	Buildah          = "buildah"
)

type ImageComparer struct {
	imageName string
	Utility   UtilType
	imagePath string
}

func (ic *ImageComparer) GetImageSize() {

}

func (ic *ImageComparer) TestBuildingImage() CompareData {
	utility := string(ic.Utility)

	// Start systemd services
	if ic.Utility == "docker" || ic.Utility == "podman" {
		utils.CallSystemctl("start", utility)
	}

	// Call building command
	command := "build"
	if utility == "buildah" {
		command = "bud"
	}

	cmd := exec.Command(utility, command, ic.imagePath, "-t", ic.imageName)

	t1 := time.Now()

	// Start executing command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command: ", err)
		os.Exit(1)
	}

	// Get ram usage of process
	kb := marks.GetCurrentMemoryUsage(cmd)

	// Wait for the command to finish and print any errors
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error: ", err)
		os.Exit(1)
	}

	t2 := time.Now()

	// Get time for executing command
	timer := marks.CreateTimeMark(t1, t2)
	ms := timer.TakeDiff()

	// Stop systemd services
	if ic.Utility == "docker" || ic.Utility == "podman" {
		utils.CallSystemctl("stop", utility+".socket")
		utils.CallSystemctl("stop", utility)
	}

	return CompareData{buildImageSize: 0,
		buildRamUsage: kb,
		buildTime:     ms}
}

func CreateImageComparer(utility UtilType, imageName string, imagePath string) *ImageComparer {
	return &ImageComparer{Utility: utility,
		imagePath: imagePath,
		imageName: imageName}
}
