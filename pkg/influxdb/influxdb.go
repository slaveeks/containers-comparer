package influxdb

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

func ConnectToInfluxDB(url string, token string) api.WriteAPIBlocking {
	client := influxdb2.NewClient(url, token)

	fmt.Println("Connected to InfluxDB!")

	writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")

	return writeAPI
}
