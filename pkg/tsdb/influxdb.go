package influxdb

import (
	"context"
	"log"
	"fmt"
	"path/filepath"
	"time"

	toml "github.com/pelletier/go-toml"
	"github.com/influxdata/influxdb-client-go/v2"
)

type TSDB struct {
    host string
    port int64
    token string
    uri string
    org string
    bucket string

}

func handleFatalError(msg string, err error){
  if err != nil {
    log.Fatal(msg, err.Error())
  }
}

func handleNonFatalError(msg string, err error){
	if err != nil {
		log.Println(err.Error())
	}
}


type SimSession struct {
    location string
    driver string
    gender string
}

func (tsdb *TSDB) Initialize(configPath string) {
    log.Println("Initializing TDSB...")
    config_path, err := filepath.Abs(configPath)
    handleFatalError("When generating config_path", err)
    config, err := toml.LoadFile(config_path)
    handleFatalError("When reading config file", err)
    tsdb.host = config.Get("influxdb_connection.host").(string)
    tsdb.port = config.Get("influxdb_connection.port").(int64)
    tsdb.uri = fmt.Sprintf("http://%s:%d", tsdb.host, tsdb.port)
    tsdb.token = config.Get("influxdb_connection.token").(string)
    tsdb.org = config.Get("influxdb_connection.org").(string)
    tsdb.bucket = config.Get("influxdb_connection.bucket").(string)
    log.Println(tsdb.uri, tsdb.org, tsdb.bucket)
}

func (ss *SimSession) Initialize(configPath string) {
    log.Println("Initializing SimSession...")
    session_path, err := filepath.Abs(configPath)
    if err != nil{
        log.Fatal(err)
    }
    config, err := toml.LoadFile(session_path)
    ss.location = config.Get("session.location").(string)

    ss.driver = config.Get("session.driver").(string)
    ss.gender = config.Get("session.gender").(string)
    log.Println("Session initialized as:", ss)
}

func (tsdb *TSDB) SaveIntDataPoint(ss *SimSession, fieldName string, value int32) {
    influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
    writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
    q := influxdb2.NewPointWithMeasurement(ss.driver).
            AddTag("location", ss.location).
            AddField(fieldName, value).
            SetTime(time.Now())
    err := writeAPI.WritePoint(context.Background(), q)
    handleNonFatalError("When saving integer datapoint", err)
    influx_client.Close()
}

func (tsdb *TSDB) SaveIntDataPointWithUnits(ss *SimSession, units string, fieldName string, value int32) {
        influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
        writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
        p := influxdb2.NewPointWithMeasurement(ss.driver).
                AddTag("location", ss.location).
                AddTag("units,", units).
                AddField(fieldName, value).
                SetTime(time.Now())
        err := writeAPI.WritePoint(context.Background(), p)
        handleNonFatalError("When saving integer datapoint with units", err)
        influx_client.Close()
}

func (tsdb *TSDB) SaveFloatDataPointWithUnits(ss *SimSession, units string, fieldName string, value float64) {
        influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
        writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
        p := influxdb2.NewPointWithMeasurement(ss.driver).
                AddTag("location", ss.location).
                AddTag("units,", units).
                AddField(fieldName, value).
                SetTime(time.Now())
        err := writeAPI.WritePoint(context.Background(), p)
        handleNonFatalError("When saving float datapoint with units", err)
        influx_client.Close()
}

func (tsdb *TSDB) SaveFloatDataPoint(ss *SimSession, fieldName string, value float64) {
        influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
        writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
        q := influxdb2.NewPointWithMeasurement(ss.driver).
                AddTag("location", ss.location).
                AddField(fieldName, value).
                SetTime(time.Now())
        err := writeAPI.WritePoint(context.Background(), q)
        handleNonFatalError("When saving float datapoint with units", err)
        influx_client.Close()
}

