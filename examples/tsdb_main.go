package main

import (
    "github.com/forzaclaudio/driving-data-engineering/pipeline/tsbd"
)

func main(){
    myTSDB := &influxdb.TSDB{}
    myTSDB.Initialize("../config/config.toml")
    mySession := &influxdb.SimSession{}
    mySession.Initialize("../config/session.toml")
    myTSDB.SaveIntDataPointWithUnits(mySession, "km/h", "speed", 180.0)
    myTSDB.SaveIntDataPoint(mySession, "throttle", 100.0)
