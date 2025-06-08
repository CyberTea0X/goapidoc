package tests

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/CyberTea0X/goapidoc"
	"gopkg.in/yaml.v3"
)

func TestDocumentFieldOrder(t *testing.T) {
	doc := &goapidoc.Document{
		OpenApiVersion: "3.0.0",
		Info: goapidoc.Info{
			Title:       "Test API",
			Description: "API with ordered fields",
			Version:     "1.0.0",
		},
		Tags: []goapidoc.Tag{
			{Name: "users", Description: "User operations"},
		},
		Paths: map[string]goapidoc.Path{
			"/users": {
				Get: &goapidoc.Method{
					Summary:     "Get users",
					OperationId: "getUsers",
					Responses: map[string]goapidoc.Response{
						"200": {
							Description: "OK",
						},
					},
				},
			},
		},
		Components: &goapidoc.Components{
			Schemas: goapidoc.Schemas{
				"User": {
					Type: goapidoc.Object,
					Properties: map[string]interface{}{
						"id":   goapidoc.SchemaInt32,
						"name": goapidoc.SchemaString,
					},
				},
			},
		},
		Servers: []goapidoc.Server{
			{Url: "http://localhost:8080"},
		},
		Security: []goapidoc.Security{
			{"oauth2": {"read", "write"}},
		},
	}

	// === Проверка порядка в JSON ===
	jsonData, err := json.Marshal(doc)
	if err != nil {
		t.Errorf("failed to marshal JSON: %v", err)
		return
	}
	jsonStr := string(jsonData)

	expectedOrder := []string{
		"openapi",
		"servers",
		"info",
		"security",
		"tags",
		"paths",
		"components",
	}

	if !isOrdered(jsonStr, expectedOrder, t) {
		t.Errorf("JSON fields are not in expected order\nExpected: %v\nGot: %s", expectedOrder, jsonStr)
	}

	// === Проверка порядка в YAML ===

	// небольшой костыль для yaml
	expectedOrder[3] = "security:\n"

	yamlBytes, err := yaml.Marshal(doc)
	if err != nil {
		t.Errorf("failed to marshal YAML: %v", err)
		return
	}
	yamlStr := string(yamlBytes)

	if !isOrdered(yamlStr, expectedOrder, t) {
		t.Errorf("YAML fields are not in expected order\nExpected: %v\nGot: %s", expectedOrder, yamlStr)
	}
}

// isOrdered проверяет, что ключи в строке идут в указанном порядке
func isOrdered(data string, expectedOrder []string, t *testing.T) bool {
	for i := range len(expectedOrder) - 1 {
		firstKey := expectedOrder[i]
		nextKey := expectedOrder[i+1]

		firstPos := strings.Index(data, firstKey)
		nextPos := strings.Index(data, nextKey)

		fmt.Printf("%s → %d | %s → %d\n", firstKey, firstPos, nextKey, nextPos)

		if firstPos == -1 || nextPos == -1 {
			return false
		}
		if firstPos > nextPos {
			t.Logf("Order violation: %s should be before %s", firstKey, nextKey)
			return false
		}
	}
	return true
}
