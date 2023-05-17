package obs_client

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/mock"
)

type VideoRecorder struct {
	obs_host       string
	obs_port       int64
	obs_pwd        string
	obs_connection string
}

type publicOBS interface {
	GetVersion()
	StartRecording()
	StopRecording()
}

type mockVideoRecorder struct{ mock.Mock }

func newMockVideoRecorder() *mockVideoRecorder {
	return &mockVideoRecorder{}
}

func (vr *VideoRecorder) Identify() string {
	id := "VideoRecorder"
	return id
}

func (vr *VideoRecorder) getParam(s string) interface{} {
	return nil
}

func (vr *VideoRecorder) ParseValues(config *toml.Tree) (err error) {
	defer func() error {
		if recover() != nil {
			err = errors.New("[VideoRecorder Initialize]: failed to parse one or various required values.")
		}
		return err
	}()
	vr.obs_host = config.Get("obs_connection.host").(string)
	vr.obs_port = config.Get("obs_connection.port").(int64)
	vr.obs_pwd = config.Get("obs_connection.password").(string)
	return nil
}

func (vr *VideoRecorder) Initialize(configPath string) error {
	log.Println("Initializing VideoRecorder...")
	config_path, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("[VideoRecorder Initialize]: failed to generate configPath: %w", err)
	}
	config, err := toml.LoadFile(config_path)
	if err != nil {
		return fmt.Errorf("[VideoRecorder Initialize]: failed to load config file: %w", err)
	}
	err = vr.ParseValues(config)
	if err != nil {
		return fmt.Errorf("[VideoRecorder Initialize]: failed to parse values: %w", err)
	}
	vr.obs_connection = fmt.Sprintf("%s:%s", vr.obs_host, strconv.Itoa(int(vr.obs_port)))
	log.Println("Will connect to the following Websocket Server:", vr.obs_connection)
	return nil
}

func (mockVR *mockVideoRecorder) GetVersion() error {
	args := mockVR.Called()
	return args.Error(0)
}

func (vr *VideoRecorder) GetVersion() error {
	log.Println("[VideoRecorder GetVersion]: Not yet implemented")
	return nil
}

func (mockVR *mockVideoRecorder) StartRecording() error {
	args := mockVR.Called()
	return args.Error(0)
}

func (vr *VideoRecorder) StartRecording() error {
	log.Println("[VideoRecorder GetVersion]: Not yet implemented")
	return nil
}

func (mockVR *mockVideoRecorder) StopRecording() error {
	args := mockVR.Called()
	return args.Error(0)
}

func (vr *VideoRecorder) StopRecording() error {
	log.Println("[VideoRecorder GetVersion]: Not yet implemented")
	return nil
}
