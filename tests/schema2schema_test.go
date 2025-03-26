package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CyberTea0X/goapidoc"
)

func TestSchemaFromSchema(t *testing.T) {
	schema := goapidoc.Schema{
		Type: "object",
		Properties: map[string]any{
			"foo": goapidoc.Schema{Type: "string"},
			"bar": goapidoc.Schema{Type: "number"},
		},
	}

	result, err := goapidoc.SchemaFromOrErr(schema)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, schema) {
		t.Errorf("expected %v, got %v", schema, result)
	}
	rawJson, err := json.Marshal(result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(rawJson) != "{\"type\":\"object\",\"properties\":{\"bar\":{\"type\":\"number\"},\"foo\":{\"type\":\"string\"}}}" {
		t.Errorf("expected %v, got %v", "{\"type\":\"object\",\"properties\":{\"bar\":{\"type\":\"number\"},\"foo\":{\"type\":\"string\"}}}", string(rawJson))
	}
}
