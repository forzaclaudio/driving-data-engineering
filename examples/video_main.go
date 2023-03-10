package video

import (
        "time"
)

func main(){
    myRecorder := &VideoRecorder{}
    myRecorder.Initialize("../../config/config.toml")
    myRecorder.GetInfo()
    myRecorder.StartRecording()
    time.Sleep(10 * time.Second)
    myRecorder.StopRecording()
}

