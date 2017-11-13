package server

import (
	//"bytes"
	//"io/ioutil"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.APIConfig.Listen != "10010"{
		t.Errorf("%q the ", config.APIConfig.Listen)
	}
}