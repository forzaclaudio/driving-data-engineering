package main

func main(){
    myTSDB := &TSDB{}
    myTSDB.Initialize()
    mySession := &SimSession{}
    mySession.Initialize()
    myTSDB.SaveDataPointWithUnits(mySession, "km/h", "speed", 180.0)
    myTSDB.SaveDataPoint(mySession, "throttle", 100.0)
}
