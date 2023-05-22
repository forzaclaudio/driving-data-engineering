package f1_telemetry

import (
	"errors"
	"strings"
	"testing"

	toml "github.com/pelletier/go-toml"
)

func setupTestConfigurations() map[int]*toml.Tree {
	telemetry_conf := make(map[int]*toml.Tree)
	telemetry_conf[1], _ = toml.Load(`[telemetry_client]`)
	telemetry_conf[2], _ = toml.Load(`[telemetry_client]
udp_port = 20777`)
	return telemetry_conf
}

func Test_TelemetryClient_Initialization_NonExistingFile(t *testing.T) {
	want := errors.New("[Telemetry Client Initialize]: failed to load config file:")
	tc := new(TelemetryClient)
	got := tc.Initialize("non-existing-config")
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
}

func Test_TelemetryClient_ParseValues(t *testing.T) {
	tc := new(TelemetryClient)
	configurations := setupTestConfigurations()
	want := errors.New("[Telemetry Client Initialize]: failed to parse one or various required values.")
	got := tc.ParseValues(configurations[1])
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
	got = tc.ParseValues(configurations[2])
	if got != nil {
		t.Errorf("Unexpected error when parsing config. %s", got.Error())
	}
}

func Test_TelemetryClient(t *testing.T) {
	want := &TelemetryClient{20777}
	got := &TelemetryClient{}
	got.Initialize("conf_test.toml")
	if got.udp_port != want.udp_port {
		t.Errorf("got.udp_port = %q, want.upd_port = %q", got.udp_port, want.udp_port)
	}
}

func Test_TelemetryClient_GetParams(t *testing.T) {
	got := &TelemetryClient{20777}
	if got.GetParams() == nil {
		t.Errorf("Unexpected value returned by method GetParams")
	}
}
