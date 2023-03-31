package comparer

import (
	"awesomeProject/src/comparer/images"
	"fmt"
	"strconv"
)

type UtilTester struct {
	name       images.UtilType
	path       string
	testNumber int
}

func (u *UtilTester) BuildImage() {

	imageComparer := images.CreateImageComparer(u.name, "image"+strconv.Itoa(u.testNumber), u.path)

	data := imageComparer.TestBuildingImage()

	fmt.Println(data)

}

func CreateUtilTester(name images.UtilType, path string, testNumber int) *UtilTester {
	return &UtilTester{name: name,
		path:       path,
		testNumber: testNumber}
}
