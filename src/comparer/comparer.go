package comparer

import (
	"awesomeProject/src/marks"
	"log"
	"os/exec"
	"time"
)

type UtilType string

const (
	Podman UtilType = "podman"
	Docker          = "docker"
)

type UtilTester struct {
	name UtilType
	path string
}

func (u *UtilTester) BuildImage() {
	cmd := exec.Command(string(u.name), "build", u.path)

	t1 := time.Now()

	marks.GetCurrentMemoryUsage()

	err := cmd.Run()
	marks.GetCurrentMemoryUsage()

	if err != nil {
		log.Fatal(err)
	}
	t2 := time.Now()

	timer := marks.CreateTimeMark(t1, t2)

	timer.TakeDiff()
}

func CreateUtilTester(name UtilType, path string) *UtilTester {
	return &UtilTester{name: name,
		path: path}
}
