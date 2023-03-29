package db

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

type Data struct {
	Time        int
	Ram         int
	Name        string
	CheckNumber int
}

type Db struct {
	api api.WriteAPIBlocking
}

func (db *Db) Insert(data Data) {
	p := influxdb2.NewPointWithMeasurement(data.Name).
		AddField("duration", data.Time).
		AddField("ram", data.Ram).
		AddField("check_number", data.CheckNumber)
	err := db.api.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted!")
}

func CreateDb(api api.WriteAPIBlocking) *Db {
	return &Db{api: api}
}
