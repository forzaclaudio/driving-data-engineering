package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	toml "github.com/pelletier/go-toml"
)

type DataStorer interface {
	Initialize()
	SaveDataPoint()
}
type SessionConfigurer interface {
	Initialize()
	SaveDataPointWithUnits()
	SaveDataPoint()
}

type TSDB struct {
	host  string
	port  int64
	token string
	uri   string
}

type SimSession struct {
	location string
	org      string
	bucket   string
	driver   string
	gender   string
}

func (tsdb *TSDB) Initialize() {
	log.Println("Initializing TDSB...")
	config_path, err := filepath.Abs("../../config/config.toml")
	if err != nil {
		log.Println("When generating config_path")
		log.Fatal(err)
	}
	config, err := toml.LoadFile(config_path)
	if err != nil {
		log.Println("When reading config file")
		log.Fatal(err)
	}
	tsdb.uri = fmt.Sprintf("http://%s:%d", config.Get("influxdb_connection.host").(string), config.Get("influxdb_connection.port").(int64))
	tsdb.token = config.Get("influxdb_connection.token").(string)
	log.Println(tsdb.uri)
}

func (ss *SimSession) Initialize() {
	log.Println("Initializing SimSession...")
	session_path, err := filepath.Abs("../../config/session.toml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := toml.LoadFile(session_path)
	ss.location = config.Get("session.location").(string)
	ss.org = config.Get("session.org").(string)
	ss.bucket = config.Get("session.bucket").(string)
	ss.driver = config.Get("session.driver").(string)
	ss.gender = config.Get("session.gender").(string)
	log.Println("Session initialized as:", ss)
}

func (tsdb *TSDB) SaveFloatDataPointWithUnits(ss *SimSession, units string, fieldName string, value float64) {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(ss.org, ss.bucket)
	p := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddTag("units,", units).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
	influx_client.Close()
}

func (tsdb *TSDB) SaveIntDataPointWithUnits(ss *SimSession, units string, fieldName string, value int32) {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(ss.org, ss.bucket)
	p := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddTag("units,", units).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
	influx_client.Close()
}

func (tsdb *TSDB) SaveFloatDataPoint(ss *SimSession, fieldName string, value float64) {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(ss.org, ss.bucket)
	q := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), q)
	if err != nil {
		panic(err)
	}
	influx_client.Close()
}

func (tsdb *TSDB) SaveIntDataPoint(ss *SimSession, fieldName string, value int32) {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(ss.org, ss.bucket)
	q := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), q)
	if err != nil {
		panic(err)
	}
	influx_client.Close()
}
