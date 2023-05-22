package obs_client

import (
	"errors"
	"strings"
	"testing"

	"github.com/andreykaipov/goobs"
	toml "github.com/pelletier/go-toml"
)

type MockGoobs struct {
	client *goobs.Client
}

type MockGeneral struct {
	subclients mockSubclients
}

type mockSubclients struct {
	General *goobs.Client
}

func setupTestConfigurations() map[int]*toml.Tree {
	obs_conf := make(map[int]*toml.Tree)
	obs_conf[1], _ = toml.Load(`[obs_connection]`)
	obs_conf[2], _ = toml.Load(`[obs_connection]
host = "some-host"
`)
	obs_conf[3], _ = toml.Load(`[obs_connection]
host = "some-host"
port = 1234
`)
	obs_conf[4], _ = toml.Load(`[obs_connection]
    host = "some-host"
    port = 1234
    password = "some-password"
`)
	return obs_conf
}

func Test_VideoRecorder_InitializationNonExistingFile(t *testing.T) {
	want := errors.New("[VideoRecorder Initialize]: failed to load config file:")
	testVR := new(VideoRecorder)
	got := testVR.Initialize("non-existing-config")
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
}

func Test_VideoRecorder_Initialize(t *testing.T) {
	want := &VideoRecorder{"some-host", 1234, "some-password", "some-connection"}
	got := &VideoRecorder{}
	got.Initialize("conf_test.toml")
	if got.obs_host != want.obs_host {
		t.Errorf("got.obs_host = %q, want.obs_host = %q", got.obs_host, want.obs_host)
	}
	if got.obs_port != want.obs_port {
		t.Errorf("got.obs_port = %q, want.obs_port = %q", got.obs_port, want.obs_port)
	}
	if got.obs_pwd != want.obs_pwd {
		t.Errorf("got.obs_pwd = %q, want.obs_pwd = %q", got.obs_pwd, want.obs_pwd)
	}
}

func Test_VideoRecorder_ParseValues(t *testing.T) {
	testVR := new(VideoRecorder)
	configurations := setupTestConfigurations()
	want := errors.New("[VideoRecorder Initialize]: failed to parse one or various required values.")
	for i := 1; i <= 3; i++ {
		got := testVR.ParseValues(configurations[i])
		if !strings.Contains(got.Error(), want.Error()) {
			t.Errorf("got = %q want = %q", got.Error(), want.Error())
		}
	}
	got := testVR.ParseValues(configurations[4])
	if got != nil {
		t.Errorf("Unexpected error when parsing config. %s", got.Error())
	}
}

func Test_GetVersion(t *testing.T) {
	vr := new(VideoRecorder)
	vr.Initialize("conf_test.toml")
	mockVR := newMockVideoRecorder()
	mockVR.On("GetVersion").Return(nil)
	got := mockVR.GetVersion()
	if got != nil {
		t.Errorf("Unexpected error when parsing calling GetVersion")
	}
}

func Test_StartRecording(t *testing.T) {
	vr := new(VideoRecorder)
	vr.Initialize("conf_test.toml")
	mockVR := newMockVideoRecorder()
	mockVR.On("StartRecording").Return(nil)
	got := mockVR.StartRecording()
	if got != nil {
		t.Errorf("Unexpected error when parsing calling StartRecording")
	}
}

func Test_StopRecording(t *testing.T) {
	vr := new(VideoRecorder)
	vr.Initialize("conf_test.toml")
	mockVR := newMockVideoRecorder()
	mockVR.On("StartRecording").Return(nil)
	got := mockVR.StartRecording()
	if got != nil {
		t.Errorf("Unexpected error when parsing calling StartRecording")
	}
}

func Test_VideoRecorder_GetParams(t *testing.T) {
	got := VideoRecorder{"some-host", 0, "some-password", "some-connection"}
	if got.GetParams() == nil {
		t.Errorf("Unexpected value returned by method GetParams")
	}
}
