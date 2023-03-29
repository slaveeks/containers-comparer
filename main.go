package main

import (
	"awesomeProject/pkg/influxdb"
	"awesomeProject/src/comparer"
	"awesomeProject/src/db"
	"strconv"
	"time"
)

func main() {
	api := influxdb.ConnectToInfluxDB("localhost:8086", "user:password")
	d := db.CreateDb(api)
	for i := 0; i <= 5; i++ {
		c := comparer.CreateUtilTester("podman", "./test-app/check"+strconv.Itoa(i), d, i)
		c.BuildImage()
		time.Sleep(20 * time.Second)
		c.RunContainer()
		time.Sleep(20 * time.Second)
	}
}
