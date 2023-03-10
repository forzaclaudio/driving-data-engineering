package influxdb

import (
	"testing"
)

func TestInfluxdb(t *testing.T){
    want := &TSDB{"localhost", 1234, "<some-token>", "http://localhost:1234"}
    got := &TSDB{}
    got.Initialize("conf_test.toml")
    if got.host != want.host{
        t.Errorf("got.host = %q, want.host = %q", got.host, want.host)
    }
    if got.port != want.port{
        t.Errorf("got.port = %q, want.port = %q", got.port, want.port)
    }
    if got.uri != want.uri{
        t.Errorf("got.uri = %q, want.uri = %q", got.uri, want.uri)
    }
    if got.token != want.token{
        t.Errorf("got.token = %q, want.token = %q", got.token, want.token)
    }
}
