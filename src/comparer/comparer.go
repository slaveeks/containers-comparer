package comparer

import (
	"awesomeProject/src/db"
	"awesomeProject/src/marks"
	"fmt"
	"os"
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
	db   *db.Db
}

func (u *UtilTester) BuildImage() {
	cmd := exec.Command(string(u.name), "build", u.path)

	t1 := time.Now()

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command: ", err)
		os.Exit(1)
	}

	marks.GetCurrentMemoryUsage(cmd)

	// Wait for the command to finish and print any errors
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error: ", err)
		os.Exit(1)
	}

	t2 := time.Now()

	timer := marks.CreateTimeMark(t1, t2)

	timer.TakeDiff()
}

func CreateUtilTester(name UtilType, path string, db *db.Db) *UtilTester {
	return &UtilTester{name: name,
		path: path,
		db:   db}
}
