package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/CyberTea0X/goapidoc"
)

func TestCallbackSpecGeneration(t *testing.T) {
	// Создаем документ OpenAPI с коллбэком
	doc := &goapidoc.Document{
		OpenApiVersion: "3.0.4",
		Info: goapidoc.Info{
			Title:       "test",
			Description: "API with callback definition",
			Version:     "0.0.0",
		},
		Paths: map[string]goapidoc.Path{
			"/subscribe": {
				Post: &goapidoc.Method{
					Summary: "Subscribe to a webhook",
					RequestBody: goapidoc.RequestWithJson("", goapidoc.Schema{
						Type: "object",
						Properties: map[string]any{
							"callbackUrl": map[string]any{
								"type":    "string",
								"format":  "uri",
								"example": "https://myserver.com/send/callback/here",
							},
						},
						Required: []string{"callbackUrl"},
					}, true),
					Callbacks: map[string]goapidoc.Paths{
						"myEvent": {
							"{$request.body#/callbackUrl}": {
								Post: &goapidoc.Method{
									RequestBody: goapidoc.RequestWithJson("", goapidoc.Schema{
										Type: "object",
										Properties: map[string]any{
											"message": map[string]any{
												"type":    "string",
												"example": "Some event happened",
											},
										},
										Required: []string{"message"},
									}, true),
									Responses: map[string]goapidoc.Response{
										"200": {
											Description: "Your server returns this code if it accepts the callback",
										},
									},
								},
							},
						},
					},
					Responses: map[string]goapidoc.Response{
						"201": {
							Description: "Webhook created",
						},
					},
				},
			},
		},
	}

	// Проверяем JSON-сериализацию
	jsonData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	jsonStr := string(jsonData)

	// Проверяем ключевые элементы в JSON
	checkStrings := []string{
		"callbacks",
		"myEvent",
		"{$request.body#/callbackUrl}",
		"\"post\": {",
		"message",
		"200",
		"Webhook created",
	}

	for _, substr := range checkStrings {
		if !strings.Contains(jsonStr, substr) {
			t.Errorf("JSON is missing expected substring: %s\nFull JSON: %s", substr, jsonStr)
		}
	}
}
