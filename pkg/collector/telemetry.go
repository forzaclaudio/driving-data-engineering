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
fmt.Println(client)*/
