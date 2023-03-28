package main

import (
	"awesomeProject/pkg/influxdb"
	"awesomeProject/src/comparer"
)

func main() {
	influxdb.ConnectToInfluxDB("localhost:8086", "user:password")
	c := comparer.CreateUtilTester("podman", "./test-app/check")
	c.BuildImage()
}
