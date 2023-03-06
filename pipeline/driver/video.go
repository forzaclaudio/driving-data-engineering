package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	toml "github.com/pelletier/go-toml"
	"github.com/andreykaipov/goobs"
)

type VideoCapturer interface{
    Initialize()
    GetInfo()
    StartRecording()
    StopRecording()
}

type VideoRecorder struct {
    obs_host string
    obs_port int64
    obs_pwd string
    obs_connection string
}

func (x *VideoRecorder) StartRecording() {
    obs_client, err := goobs.New(x.obs_connection, goobs.WithPassword(x.obs_pwd))
    if err != nil{
        log.Fatal("When getting OBS client", err)
    }
    defer obs_client.Disconnect()
    _, err =  obs_client.Record.StartRecord()
    if err != nil{
       log.Fatal("On StartRecording", err)
    }
}

func (x *VideoRecorder) StopRecording() {
    obs_client, err := goobs.New(x.obs_connection, goobs.WithPassword(x.obs_pwd))
    if err != nil{
        log.Fatal("When getting OBS client", err)
    }
    defer obs_client.Disconnect()
    outFilePath, err := obs_client.Record.StopRecord()
    if err != nil{
       log.Fatal("On StopRecording", err)
    }
    log.Println("Recording saved to file:", outFilePath)
}

func (x *VideoRecorder) GetInfo() {
    obs_client, err := goobs.New(x.obs_connection, goobs.WithPassword(x.obs_pwd))
    if err != nil{
        log.Fatal("When getting OBS client", err)
    }
    defer obs_client.Disconnect()
    version, err :=  obs_client.General.GetVersion()
    if err != nil{
        log.Fatal("On GetInfo", err)
    }
    log.Println("OBS studio version:", version.ObsVersion, "Winsocks Server version:", version.ObsWebSocketVersion)
}

func (x *VideoRecorder) Initialize() {
    log.Println("Initializing VideoRecorder...")
    config_path, err := filepath.Abs("../../config/config.toml")
    if err != nil{
        log.Fatal(err)
    }
    config, err := toml.LoadFile(config_path)
    if err != nil{
        log.Fatal(err)
    } else {
	    x.obs_host = config.Get("obs_connection.host").(string)
	    x.obs_port = config.Get("obs_connection.port").(int64)
	    x.obs_pwd = config.Get("obs_connection.password").(string)
	    log.Println("Will connect to the following Websocket Server:", x.obs_host, x.obs_port)
    }
    x.obs_connection = fmt.Sprintf("%s:%s", x.obs_host, strconv.Itoa(int(x.obs_port)))
}

