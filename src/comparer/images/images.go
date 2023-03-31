package images

import (
	"awesomeProject/src/marks"
	"awesomeProject/src/utils"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type UtilType string

type CompareData struct {
	buildTime      int
	buildRamUsage  int
	buildImageSize string
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

	cmd := exec.Command(utility, command, "-t", ic.imageName, ic.imagePath)

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

	cmd = exec.Command(utility, "images", ic.imageName)
	var imagesOutput bytes.Buffer
	cmd.Stdout = &imagesOutput
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// extract the size of the image from the `docker images` output
	size := extractImageSize(imagesOutput.String())

	// Stop systemd services
	if ic.Utility == "docker" || ic.Utility == "podman" {
		utils.CallSystemctl("stop", utility+".socket")
		utils.CallSystemctl("stop", utility)
	}

	return CompareData{buildImageSize: size,
		buildRamUsage: kb,
		buildTime:     ms}
}

func extractImageSize(output string) string {
	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		return ""
	}
	parts := strings.Fields(lines[1])
	if len(parts) < 8 {
		return ""
	}
	return parts[6]
}

func CreateImageComparer(utility UtilType, imageName string, imagePath string) *ImageComparer {
	return &ImageComparer{Utility: utility,
		imagePath: imagePath,
		imageName: imageName}
}
