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
}

func handleFatalError(msg string, err error){
  if err != nil {
    log.Fatal(msg, err.Error())
  }
}

func handleNonFatalError(err error){
	if err != nil {
		log.Println(err.Error())
	}
}


type SimSession struct {
    location string
    org string
    bucket string
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
    log.Println(tsdb.uri)
}

func (ss *SimSession) Initialize() {
    log.Println("Initializing SimSession...")
    session_path, err := filepath.Abs("../../config/session.toml")
    if err != nil{
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

func (tsdb *TSDB) SaveDataPointWithUnits(ss *SimSession, units string, fieldName string, value float64) {
    influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
    writeAPI := influx_client.WriteAPIBlocking("volvo-data-engineering", "f1-sim")
    p := influxdb2.NewPointWithMeasurement(ss.driver).
        AddTag("location", ss.location).
	AddTag("units,", units).
	AddField(fieldName, value).
	SetTime(time.Now())
    err := writeAPI.WritePoint(context.Background(),p)
                if err != nil{
                    panic(err)
                }
    influx_client.Close()
}

func (tsdb *TSDB) SaveDataPoint(ss *SimSession, fieldName string, value float64) {
    influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
    writeAPI := influx_client.WriteAPIBlocking("volvo-data-engineering", "f1-sim")
    q := influxdb2.NewPointWithMeasurement(ss.driver).
        AddTag("location", ss.location).
	AddField(fieldName, value).
	SetTime(time.Now())
    err := writeAPI.WritePoint(context.Background(), q)
                if err != nil{
                    panic(err)
                }
    influx_client.Close()
}
