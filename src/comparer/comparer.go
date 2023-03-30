package comparer

import (
	"awesomeProject/src/db"
	"awesomeProject/src/marks"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UtilType string

const (
	Podman UtilType = "podman"
	Docker          = "docker"
)

type UtilTester struct {
	name       UtilType
	path       string
	db         *db.Db
	testNumber int
}

func (u *UtilTester) BuildImage() {
	err := u.CallSystemctl("start")

	if err != nil {
		fmt.Println("Error systemctl start: ", err)
		os.Exit(1)
	}

	cmd := exec.Command(string(u.name), "build", u.path, "-t", "image"+strconv.Itoa(u.testNumber))

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
		u.testNumber,
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

	if command == "stop" {
		cmd = exec.Command("systemctl", command, string(u.name)+".socket")

		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	cmd = exec.Command("systemctl", command, string(u.name))

	return cmd.Run()
}

func (u *UtilTester) RunContainer() {
	err := u.CallSystemctl("start")

	if err != nil {
		fmt.Println("Error systemctl start: ", err)
		os.Exit(1)
	}

	t1 := time.Now()
	cmd := exec.Command(string(u.name), "run", "-d", "-t", "image"+strconv.Itoa(u.testNumber))
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

	ramCmd := exec.Command(string(u.name), "stats", "--no-stream", containerID, "--format", "{{.MemUsage}}")
	ramOut, err := ramCmd.Output()

	if err != nil {
		fmt.Println("Error while getting docker stats", err)
		return
	}

	ramUsage := strings.TrimSpace(string(ramOut))

	re := regexp.MustCompile(`^\d+`)

	match := re.FindString(ramUsage)

	num, _ := strconv.Atoi(match)

	fmt.Printf("Container RAM usage: %d\n", num)

	if err != nil {
		fmt.Println("Error systemctl stop: ", err)
		os.Exit(1)
	}
}

func CreateUtilTester(name UtilType, path string, db *db.Db, testNumber int) *UtilTester {
	return &UtilTester{name: name,
		path:       path,
		db:         db,
		testNumber: testNumber}
}
