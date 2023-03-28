package influxdb

import influxdb2 "github.com/influxdata/influxdb-client-go"

func ConnectToInfluxDB(url string, token string) influxdb2.Client {
	return influxdb2.NewClient(url, token)
}
