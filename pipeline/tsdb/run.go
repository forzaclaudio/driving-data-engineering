package main

import (
    "github.com/forzaclaudio/driving-data-engineering/pipeline/tsbd"
)

func main(){
    myTSDB := &influxdb.TSDB{}
    myTSDB.Initialize()
    mySession := &influxdb.SimSession{}
    mySession.Initialize()
    myTSDB.SaveDataPointWithUnits(mySession, "km/h", "speed", 180.0)
    myTSDB.SaveDataPoint(mySession, "throttle", 100.0)
}
