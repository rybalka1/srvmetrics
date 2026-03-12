// internal/handlers/handlers.go
package handlers

import (
	"github.com/rybalka1/srvmetrics/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

var memStorage *storage.MemStorage

func SetStorage(storage *storage.MemStorage) {
	memStorage = storage
}

func UpdateMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	metricType := pathParts[3]
	metricName := pathParts[4]

	if metricType != "gauge" && metricType != "counter" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var value interface{}
	var err error

	switch metricType {
	case "gauge":
		value, err = strconv.ParseFloat(metricName, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		memStorage.SetGauge(metricName, value.(float64))

	case "counter":
		value, err = strconv.ParseInt(metricName, 10, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		memStorage.AddCounter(metricName, value.(int64))
	}

	w.WriteHeader(http.StatusOK)
}
