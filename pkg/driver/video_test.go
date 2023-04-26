package video

import (
	"testing"
)

func TestTSDB(t *testing.T){
    want := &VideoRecorder{"some-host", 1234, "some-password", "some-connection"}
    got := &VideoRecorder{}
    got.Initialize("conf_test.toml")
    if got.obs_host != want.obs_host{
        t.Errorf("got.obs_host = %q, want.obs_host = %q", got.obs_host, want.obs_host)
    }
}
