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
	err := u.CallSystemctl("start")

	if err != nil {
		fmt.Println("Error systemctl start: ", err)
		os.Exit(1)
	}

	cmd := exec.Command(string(u.name), "build", u.path)

	t1 := time.Now()

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command: ", err)
		os.Exit(1)
	}

	kb := marks.GetCurrentMemoryUsage(cmd)

	// Wait for the command to finish and print any errors
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error: ", err)
		os.Exit(1)
	}

	t2 := time.Now()

	timer := marks.CreateTimeMark(t1, t2)

	ms := timer.TakeDiff()

	name := string(u.name)

	data := db.Data{ms,
		kb,
		name,
	}

	u.db.Insert(data)

	err = u.CallSystemctl("stop")

	if err != nil {
		fmt.Println("Error systemctl stop: ", err)
		os.Exit(1)
	}
}

func (u *UtilTester) CallSystemctl(command string) error {
	var cmd *exec.Cmd

	cmd = exec.Command("systemctl", command, string(u.name)+".socket")

	err := cmd.Run()

	if err != nil {
		return err
	}

	cmd = exec.Command("systemctl", command, string(u.name))

	return cmd.Run()
}

func CreateUtilTester(name UtilType, path string, db *db.Db) *UtilTester {
	return &UtilTester{name: name,
		path: path,
		db:   db}
}
