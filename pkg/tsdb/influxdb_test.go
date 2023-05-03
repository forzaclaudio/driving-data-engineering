package influxdb

import (
	"errors"
	"strings"
	"testing"

	toml "github.com/pelletier/go-toml"
)

func setupTestConfigurations() (map[int]*toml.Tree, map[int]*toml.Tree) {
	connection_conf := make(map[int]*toml.Tree)
	session_conf := make(map[int]*toml.Tree)
	connection_conf[1], _ = toml.Load(`[influxdb_connection]`)
	connection_conf[2], _ = toml.Load(`[influxdb_connection]
host = "localhost"`)
	connection_conf[3], _ = toml.Load(`[influxdb_connection]
host = "localhost"
port = 1234`)
	connection_conf[4], _ = toml.Load(`[influxdb_connection]
host = "localhost"
port = 1234
token = "<some-token>"`)
	connection_conf[5], _ = toml.Load(`[influxdb_connection]
host = "localhost"
port = 1234
token = "<some-token>"
org = "some-org"
`)
	connection_conf[6], _ = toml.Load(`[influxdb_connection]
host = "localhost"
port = 1234
token = "<some-token>"
org = "some-org"
bucket = "some-bucket"
`)
	session_conf[1], _ = toml.Load(`[session]`)
	session_conf[2], _ = toml.Load(`[session]
location = "some-location"
`)
	session_conf[3], _ = toml.Load(`[session]
location = "some-location"
driver = "some-driver"
`)
	session_conf[4], _ = toml.Load(`[session]
location = "some-location"
driver = "some-driver"
gender = "some-gender"
`)
	return connection_conf, session_conf
}

func TestTSDMInitializationNonExistingFile(t *testing.T) {
	want := errors.New("[TSDB Initialize]: failed to load config file:")
	testTSDB := new(TSDB)
	got := testTSDB.Initialize("non-existing-config")
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
}

func TestTSDBParseValues(t *testing.T) {
	tsdb := new(TSDB)
	configurations, _ := setupTestConfigurations()
	want := errors.New("[TSDB Initialize]: failed to parse one or various required values.")
	for i := 1; i <= 5; i++ {
		got := tsdb.ParseValues(configurations[i])
		if !strings.Contains(got.Error(), want.Error()) {
			t.Errorf("got = %q want = %q", got.Error(), want.Error())
		}
	}
	got := tsdb.ParseValues(configurations[6])
	if got != nil {
		t.Errorf("Unexpected error when parsing config. %s", got.Error())
	}
}

func TestTSDB(t *testing.T) {
	want := &TSDB{"localhost", 1234, "<some-token>", "http://localhost:1234", "some-org", "some-bucket"}
	got := &TSDB{}
	got.Initialize("conf_test.toml")
	if got.host != want.host {
		t.Errorf("got.host = %q, want.host = %q", got.host, want.host)
	}
	if got.port != want.port {
		t.Errorf("got.port = %q, want.port = %q", got.port, want.port)
	}
	if got.token != want.token {
		t.Errorf("got.token = %q, want.token = %q", got.token, want.token)
	}
	if got.org != want.org {
		t.Errorf("got.org = %q, want.org = %q", got.org, want.org)
	}
	if got.bucket != want.bucket {
		t.Errorf("got.bucket = %q, want.bucket = %q", got.bucket, want.bucket)
	}
	if got.uri != want.uri {
		t.Errorf("got.uri = %q, want.uri = %q", got.uri, want.uri)
	}
}

func TestParseSessionValues(t *testing.T) {
	test_session := new(SimSession)
	_, configurations := setupTestConfigurations()
	want := errors.New("[SimSession Initialize]: failed to parse one or various required values.")
	for i := 1; i <= 3; i++ {
		got := test_session.ParseValues(configurations[i])
		if !strings.Contains(got.Error(), want.Error()) {
			t.Errorf("got = %q want = %q", got.Error(), want.Error())
		}
	}
	got := test_session.ParseValues(configurations[4])
	if got != nil {
		t.Errorf("Unexpected error when parsing config. %s", got.Error())
	}
}

func TestSimSession(t *testing.T) {
	want := &SimSession{"some-location", "some-driver", "some-gender"}
	got := &SimSession{}
	got.Initialize("conf_test.toml")
	if got.location != want.location {
		t.Errorf("got.location = %q, want.location = %q", got.location, want.location)
	}
	if got.driver != want.driver {
		t.Errorf("got.driver = %q, want.driver = %q", got.driver, want.driver)
	}
	if got.gender != want.gender {
		t.Errorf("got.gender = %q, want.gender = %q", got.gender, want.gender)
	}
}

func TestSaveIntDataPoint(t *testing.T) {
	testTSDB := &TSDB{"localhost", 1234, "<some-token>", "http://localhost:1234",
		"some-org", "some-bucket"}
	testSession := &SimSession{"some-location", "some-driver", "some-gender"}
	want := errors.New("[TSDB SaveIntDataPoint]:")
	got := testTSDB.SaveIntDataPoint(testSession, "some-field", 0)
	if got == nil {
		t.Errorf("Expected an error but didn't get one.")
	}
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
	want = errors.New("[TSDB SaveIntDataPointWithUnits]:")
	got = testTSDB.SaveIntDataPointWithUnits(testSession, "some-units", "some-field", 0)
	if got == nil {
		t.Errorf("Expected an error but didn't get one.")
	}
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
}

func TestSaveFloatDataPoint(t *testing.T) {
	testTSDB := &TSDB{"localhost", 1234, "<some-token>", "http://localhost:1234",
		"some-org", "some-bucket"}
	testSession := &SimSession{"some-location", "some-driver", "some-gender"}
	want := errors.New("[TSDB SaveFloatDataPoint]:")
	got := testTSDB.SaveFloatDataPoint(testSession, "some-field", 0)
	if got == nil {
		t.Errorf("Expected an error but didn't get one.")
	}
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
	want = errors.New("[TSDB SaveFloatDataPointWithUnits]:")
	got = testTSDB.SaveFloatDataPointWithUnits(testSession, "some-units", "some-field", 0)
	if got == nil {
		t.Errorf("Expected an error but didn't get one.")
	}
	if !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("got = %q want = %q", got.Error(), want.Error())
	}
}
