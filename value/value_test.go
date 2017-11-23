package value

import (
	"encoding/json"
	"testing"
)

func TestRaw_Get(t *testing.T) {
	testStr := `[{"method":"get","config":{"ensure":1,"admin":true,"island":"asset-island"}}]`
	var testValue interface{}
	json.Unmarshal([]byte(testStr), &testValue)

	r := New(testValue)
	if v, err := r.Get("0"); err != nil {
		t.Error(err)
	} else if v == nil {
		t.Error("v is nil")
	}

	if v := r.String("0.method"); v != "get" {
		t.Error("need get, get ", v)
	}
	if v := r.Int("0.config.ensure"); v != 1 {
		t.Error("need 1, get", v)
	}
	if v := r.Bool("0.config.admin"); !v {
		t.Error("need true, get", v)
	}
}
