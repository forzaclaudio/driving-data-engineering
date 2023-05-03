package influxdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	toml "github.com/pelletier/go-toml"
)

type TSDB struct {
	host   string
	port   int64
	token  string
	uri    string
	org    string
	bucket string
}

func handleNonFatalError(msg string, err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

type SimSession struct {
	location string
	driver   string
	gender   string
}

func (tsdb *TSDB) ParseValues(config *toml.Tree) (err error) {
	defer func() error {
		if recover() != nil {
			err = errors.New("[TSDB Initialize]: failed to parse one or various required values.")
		}
		return err
	}()
	tsdb.host = config.Get("influxdb_connection.host").(string)
	tsdb.port = config.Get("influxdb_connection.port").(int64)
	tsdb.uri = fmt.Sprintf("http://%s:%d", tsdb.host, tsdb.port)
	tsdb.token = config.Get("influxdb_connection.token").(string)
	tsdb.org = config.Get("influxdb_connection.org").(string)
	tsdb.bucket = config.Get("influxdb_connection.bucket").(string)
	log.Println("Connection initialized as:", tsdb.uri, tsdb.org, tsdb.bucket)
	return nil
}

func (tsdb *TSDB) Initialize(configPath string) error {
	log.Println("Initializing TSDB...")
	config_path, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("[TSDB Initialize]: failed to generate configPath: %w", err)
	}
	config, err := toml.LoadFile(config_path)
	if err != nil {
		return fmt.Errorf("[TSDB Initialize]: failed to load config file: %w", err)
	}
	err = tsdb.ParseValues(config)
	if err != nil {
		return fmt.Errorf("[TSDB Initialize]: failed to parse values: %w", err)
	}
	return nil
}

func (ss *SimSession) ParseValues(config *toml.Tree) (err error) {
	defer func() error {
		if recover() != nil {
			err = errors.New("[SimSession Initialize]: failed to parse one or various required values.")
		}
		return err
	}()
	ss.location = config.Get("session.location").(string)
	ss.driver = config.Get("session.driver").(string)
	ss.gender = config.Get("session.gender").(string)
	log.Println("Session initialized as:", ss)
	return nil
}

func (ss *SimSession) Initialize(configPath string) error {
	log.Println("Initializing SimSession...")
	session_path, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("[Session Initialize]: failed to generate configPath: %w", err)
	}
	config, err := toml.LoadFile(session_path)
	if err != nil {
		return fmt.Errorf("[Session Initialize]: failed to load config file: %w", err)
	}
	err = ss.ParseValues(config)
	if err != nil {
		return fmt.Errorf("[Session Initialize]: failed to parse values: %w", err)
	}
	return nil
}

func (tsdb *TSDB) SaveIntDataPoint(ss *SimSession, fieldName string, value int32) error {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
	q := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), q)
	if err != nil {
		return fmt.Errorf("[TSDB SaveIntDataPoint]: %w", err)
	}
	influx_client.Close()
	return nil
}

func (tsdb *TSDB) SaveIntDataPointWithUnits(ss *SimSession, units string,
	fieldName string, value int32) error {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
	p := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddTag("units,", units).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return fmt.Errorf("[TSDB SaveIntDataPointWithUnits]: %w", err)
	}
	influx_client.Close()
	return nil
}

func (tsdb *TSDB) SaveFloatDataPointWithUnits(ss *SimSession, units string, fieldName string, value float64) error {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
	p := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddTag("units,", units).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	handleNonFatalError("When saving float datapoint with units", err)
	if err != nil {
		return fmt.Errorf("[TSDB SaveFloatDataPointWithUnits]: %w", err)
	}
	return nil
}

func (tsdb *TSDB) SaveFloatDataPoint(ss *SimSession, fieldName string, value float64) error {
	influx_client := influxdb2.NewClient(tsdb.uri, tsdb.token)
	writeAPI := influx_client.WriteAPIBlocking(tsdb.org, tsdb.bucket)
	q := influxdb2.NewPointWithMeasurement(ss.driver).
		AddTag("location", ss.location).
		AddField(fieldName, value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), q)
	if err != nil {
		return fmt.Errorf("[TSDB SaveFloatDataPoint]: %w", err)
	}
	influx_client.Close()
	return nil
}
