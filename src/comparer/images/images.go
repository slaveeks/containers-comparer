package images

import (
	"awesomeProject/src/marks"
	"awesomeProject/src/utils"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	if utility == "img" {
		cmd = exec.Command(utility, "ls")

	} else {
		cmd = exec.Command(utility, "images", ic.imageName)
	}

	var imagesOutput bytes.Buffer
	cmd.Stdout = &imagesOutput
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// extract the size of the image from the `docker images` output
	size := ic.extractImageSize(imagesOutput.String())

	// Stop systemd services
	if ic.Utility == "docker" || ic.Utility == "podman" {
		utils.CallSystemctl("stop", utility+".socket")
		utils.CallSystemctl("stop", utility)
	}

	return CompareData{buildImageSize: size,
		buildRamUsage: kb,
		buildTime:     ms}
}

func (ic *ImageComparer) extractImageSize(output string) int {
	lines := strings.Split(output, "\n")

	if string(ic.Utility) == "img" {
		for i := 0; i < len(lines); i++ {
			line := strings.Fields(lines[i])

			if strings.Contains(line[0], ic.imageName) {
				fmt.Println(line[1])
				return 1
			}
		}
	}
	parts := strings.Fields(lines[1])

	if string(ic.Utility) == "podman" || string(ic.Utility) == "buildah" {
		size, _ := strconv.Atoi(parts[len(parts)-2])
		return size
	}

	sizeStr := parts[len(parts)-1]

	size, _ := strconv.Atoi(sizeStr[:len(sizeStr)-2])

	return size
}

func CreateImageComparer(utility UtilType, imageName string, imagePath string) *ImageComparer {
	return &ImageComparer{Utility: utility,
		imagePath: imagePath,
		imageName: imageName}
}
