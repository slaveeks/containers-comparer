package main

import (
	"awesomeProject/pkg/influxdb"
	"awesomeProject/src/comparer"
	"awesomeProject/src/db"
)

func main() {
	api := influxdb.ConnectToInfluxDB("localhost:8086", "user:password")
	d := db.CreateDb(api)
	c := comparer.CreateUtilTester("podman", "./test-app/check", d)
	c.BuildImage()
}
