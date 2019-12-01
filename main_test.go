package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Mock struct{}

func (m *Mock) SendTelegramMessage(message string) ([]byte, error) {
	return []byte("hello world"), nil
}

func TestHealthEndpoint(t *testing.T) {
	server := RestController{&Mock{}}
	req, _ := http.NewRequest("GET", "/health", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HealthHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "text/plain", rr.Header().Get("Content-Type"))
	assert.Equal(t, "UP", rr.Body.String())
}

func TestSendEndpoint(t *testing.T) {
	server := RestController{&Mock{}}
	body := new(bytes.Buffer)

	payload := Request{
		Title:    "[Alerting] CPU alert",
		Mess:     "CPU load has reached 65%",
		RuleName: "CPU alert",
		RuleUrl:  "https://grafana.maslick.ru",
		State:    "alerting",
		Metrics: []EvalMatch{
			{
				Value:  70.2,
				Metric: "CPU load",
				Tags:   map[string]string{"__name__": "cpu_load"},
			},
		},
	}

	_ = json.NewEncoder(body).Encode(payload)
	req, _ := http.NewRequest("POST", "/", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.WebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, "hello world", rr.Body.String())
}
