package main

import (
        "time"
)

func main(){
    myRecorder := &VideoRecorder{}
    myRecorder.Initialize()
    myRecorder.GetInfo()
    myRecorder.StartRecording()
    time.Sleep(10 * time.Second)
    myRecorder.StopRecording()
}

