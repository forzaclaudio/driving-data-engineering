package influxdb

import (
	"testing"
)

type spyClient struct {
    Calls int
}

func (s *spyClient) NewClient(){
	s.Calls++
}

func TestTSDB(t *testing.T){
    want := &TSDB{"localhost", 1234, "<some-token>", "http://localhost:1234", "some-org", "some-bucket"}
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
    if got.org != want.org{
        t.Errorf("got.org = %q, want.org = %q", got.org, want.org)
    }
    if got.bucket != want.bucket{
        t.Errorf("got.bucket = %q, want.bucket = %q", got.bucket, want.bucket)
    }
}

func TestSimSession(t *testing.T){
    want := &SimSession{"some-location", "some-driver", "some-gender"}
    got := &SimSession{}
    got.Initialize("conf_test.toml")
    if got.location != want.location{
        t.Errorf("got.location = %q, want.location = %q", got.location, want.location)
    }
}
