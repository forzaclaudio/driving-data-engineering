package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"os/signal"
	"net/http"
	"github.com/forzaclaudio/driving-data-engineering/pkg/driver"
	"github.com/forzaclaudio/driving-data-engineering/pkg/tsdb"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/telemetry"
)

func pingUploader(vr *video.VideoRecorder, port int){
	print(vr)
	url := fmt.Sprintf("%s:%d","meh", port)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func main() {
	myTSDB := &influxdb.TSDB{}
	myTSDB.Initialize("../../config/config.toml")
	mySession := &influxdb.SimSession{}
	mySession.Initialize("../../config/session.toml")
	myRecorder := &video.VideoRecorder{}
	myRecorder.Initialize("../../config/config.toml")
	myRecorder.GetInfo()
	uploaderPort := 8080

	client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal("When creating telemetry client:", err)
	}



	// wait exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("Packet RecvCount: %d\n", client.Stats.RecvCount())
		log.Printf("Packet ErrCount: %d\n", client.Stats.ErrCount())
		myRecorder.StopRecording()
		pingUploader(myRecorder, uploaderPort)
		os.Exit(1)
	}()
	client.OnLapPacket(func(packet *packets.PacketLapData) {
		myRecorder.StartRecording()
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
