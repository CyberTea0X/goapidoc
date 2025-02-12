package tests

import (
	"encoding/json"
	"testing"

	"github.com/CyberTea0X/goapidoc"
)

func TestRecursive(t *testing.T) {
	type Pet struct {
		Age int
	}
	type User struct {
		Pet Pet
	}
	u := User{Pet: Pet{Age: 1}}
	schema := goapidoc.SchemaFrom(u)
	raw, err := json.Marshal(schema)
	if err != nil {
		t.Error(err)
	}
	expected := `{"type":"object","properties":{"pet":{"type":"object","properties":{"age":{"type":"integer","example":1}}}}}`
	if string(raw) != expected {
		t.Error("expected:" + expected + " got:" + string(raw))
	}
}
