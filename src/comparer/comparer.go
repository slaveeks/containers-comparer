package comparer

import (
	"awesomeProject/src/comparer/containers"
	"strconv"
)

type UtilTester struct {
	name       containers.UtilType
	path       string
	testNumber int
}

func (u *UtilTester) BuildImage() {

	c := containers.CreateContainersComparer("image"+strconv.Itoa(u.testNumber), u.name)

	c.RunContainer()

}

func CreateUtilTester(name containers.UtilType, path string, testNumber int) *UtilTester {
	return &UtilTester{name: name,
		path:       path,
		testNumber: testNumber}
}
