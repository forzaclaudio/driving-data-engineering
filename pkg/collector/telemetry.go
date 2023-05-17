package f1_telemetry

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type TelemetryClient struct {
	udp_port int64
}

func (tc *TelemetryClient) Identify() string {
	id := "TelemetryCollector"
	return id
}

func (tc *TelemetryClient) getParam(p string) interface{} {
	if p == "udp_port" {
		return tc.udp_port
	}
	return nil
}

func (tc *TelemetryClient) ParseValues(config *toml.Tree) (err error) {
	defer func() error {
		if recover() != nil {
			err = errors.New("[Telemetry Client Initialize]: failed to parse one or various required values.")
		}
		return err
	}()
	tc.udp_port = config.Get("telemetry_client.udp_port").(int64)
	return nil
}

func (tc *TelemetryClient) Initialize(configPath string) error {
	log.Println("Initializing Telemetry Client...")
	config_path, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("[Telemetry Client Initialize]: failed to generate configPath: %w", err)
	}
	config, err := toml.LoadFile(config_path)
	if err != nil {
		return fmt.Errorf("[Telemetry Client Initialize]: failed to load config file: %w", err)
	}
	err = tc.ParseValues(config)
	if err != nil {
		return fmt.Errorf("[Telemetry Client Initialize]: failed to parse values: %w", err)
	}
	log.Println("Telemetry client will listen on:", tc.udp_port)
	return nil
}

/*client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal("When creating telemetry client:", err)
	}
fmt.Println(client)

	// wait exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("Packet RecvCount: %d\n", client.Stats.RecvCount())
		log.Printf("Packet ErrCount: %d\n", client.Stats.ErrCount())
		//filePath := myRecorder.StopRecording()
		//requestFileUpload("192.168.1.2", uploaderPort, filePath)
		os.Exit(1)
	}()

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
*/
