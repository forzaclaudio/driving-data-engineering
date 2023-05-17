package main

import (
	"log"
	//"os"
	//"os/signal"

	influxdb "github.com/forzaclaudio/driving-data-engineering/pkg/tsdb"
	"github.com/forzaclaudio/f1-telemetry-go/pkg/packets"
	"github.com/forzaclaudio/f1-telemetry-go/pkg/telemetry"
)

/*func systemInit() map[string]bool{
	subsystem := make(map[string]bool)
	myTSDB := new(influxdb.TSDB)
	myTSDB.Initialize("../../config/config.toml")
	subsystem["tsdb"] = true
	return subsystem
}*/

func main() {
	/*	subsystem := systemInit()
		if subsystem["tsdb"] == true {
			log.Print("TSDB on")
		} else {
			log.Print("TSDB off")
		}
	*/
	myTSDB := new(influxdb.TSDB)
	err := myTSDB.Initialize("../../config/config.toml")
	//err := myTSDB.Initialize("../../pkg/tsdb/dummy_conf_test.toml")
	if err != nil {
		log.Fatal("[Pipeline Main]: ", err)
	}
	mySession := new(influxdb.SimSession)
	mySession.Initialize("../../config/session.toml")
	/*
		myRecorder := new(video.VideoRecorder)
		myRecorder.Initialize("../../config/config.toml")
		log.Printf("MyRecorder\n", myRecorder)
		myRecorder.GetInfo()
	log.Printf("Info collected successfully!\n", myRecorder)*/
	//uploaderPort := 8080
	//pingUploader("192.168.1.2", uploaderPort)
	isNotRecording := true

	client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal("When creating telemetry client:", err)
	}

	// wait exit signal
	/*c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("Packet RecvCount: %d\n", client.Stats.RecvCount())
		log.Printf("Packet ErrCount: %d\n", client.Stats.ErrCount())
		//filePath := myRecorder.StopRecording()
		//requestFileUpload("192.168.1.2", uploaderPort, filePath)
		os.Exit(1)
	}()*/

	client.OnLapPacket(func(packet *packets.PacketLapData) {
		if isNotRecording == true {
			//  myRecorder.StartRecording()
			isNotRecording = false
		}
		lap := packet.Self()
		myTSDB.SaveIntDataPointWithUnits(mySession, "ms", "current_lap_time", int32(lap.CurrentLapTimeInMS))
		myTSDB.SaveIntDataPointWithUnits(mySession, "ms", "last_lap_time", int32(lap.CurrentLapTimeInMS))
		myTSDB.SaveFloatDataPointWithUnits(mySession, "m", "lap_distance", float64(lap.LapDistance))
		myTSDB.SaveIntDataPoint(mySession, "is_current_lap_invalid", int32(lap.LapDistance))
		log.Printf("Current lap time: %d, Last lap time: %d, Distance: %f, Current lap invalid: %d", uint32(lap.CurrentLapTimeInMS), uint32(lap.LastLapTimeInMS), float32(lap.LapDistance), uint32(lap.CurrentLapInvalid))
	})

	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.Self()
		myTSDB.SaveIntDataPoint(mySession, "gear", int32(car.Gear))
		myTSDB.SaveFloatDataPointWithUnits(mySession, "km/h", "speed", float64(car.Speed))
		myTSDB.SaveFloatDataPointWithUnits(mySession, "rpm", "engine_rpm", float64(car.EngineRPM))
		myTSDB.SaveFloatDataPointWithUnits(mySession, "%", "throttle", float64(car.Throttle))
		myTSDB.SaveFloatDataPointWithUnits(mySession, "%", "brake", float64(car.Brake))
		log.Printf("Gear %d, Car speed %f, RPM: %f, Throttle: %f Break: %f", int32(car.Gear), float64(car.Speed), float64(car.EngineRPM), float64(car.Throttle)*100.0, float64(car.Brake)*100.0)

	})
	log.Println("F1 telemetry client running")
	client.Run()

}
