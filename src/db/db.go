package db

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"time"
)

type Data struct {
	Time int
	Ram  int
	Name string
}

type Db struct {
	api api.WriteAPIBlocking
}

func (db *Db) Insert(data Data) {
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("name", data.Name).
		AddField("time", data.Time).
		AddField("ram", data.Ram).
		SetTime(time.Now())
	err := db.api.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted!")
}

func CreateDb(api api.WriteAPIBlocking) *Db {
	return &Db{api: api}
}
