package db

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"time"
)

type Data struct {
	time int
	ram  int
	name string
}

type Db struct {
	api api.WriteAPIBlocking
}

func (db *Db) Insert(data Data) {
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("name", data.name).
		AddField("time", data.time).
		AddField("ram", data.ram).
		SetTime(time.Now())
	err := db.api.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
}

func CreateDb(api api.WriteAPIBlocking) *Db {
	return &Db{api: api}
}
