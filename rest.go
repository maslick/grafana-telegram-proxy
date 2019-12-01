package main

import (
	"encoding/json"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os"
	"strings"
)

var log = logging.MustGetLogger("grafana-telegram-proxy")

type RestController struct {
	Service ITelegramSender
}

func (this *RestController) Start() {
	http.HandleFunc("/health", this.HealthHandler)
	if useAuth() {
		http.HandleFunc("/", basicAuth(this.WebhookHandler))
	} else {
		http.HandleFunc("/", this.WebhookHandler)
	}
	fmt.Println("Starting server on port:", strings.Split(getPort(), ":")[1])
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func useAuth() bool {
	_, usernameOk := os.LookupEnv("USERNAME")
	_, passwordOk := os.LookupEnv("PASSWORD")
	return usernameOk && passwordOk
}

func (_ *RestController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("UP"))
}

func (this *RestController) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", 400)
		return
	}

	var message Request
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Info(message.Title + " / " + message.Mess)
	resp, err := this.Service.SendTelegramMessage(formatMessage(message))

	if err != nil {
		http.Error(w, "Message delivery failed: "+err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}

func formatMessage(message Request) string {
	mess := fmt.Sprintf("<b>%s</b>", message.Title)
	mess += "\n"
	for _, v := range message.Metrics {
		mess += fmt.Sprintf("<i>%s : %f</i>\n", v.Metric, v.Value)
	}
	mess += ""
	mess += message.RuleUrl
	return mess
}
