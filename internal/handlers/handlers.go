package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rybalka1/srvmetrics/internal/storage"
)

type Handler struct {
	storage *storage.MemStorage
}

func NewHandler(storage *storage.MemStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) UpdateMetric(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pieces := strings.Split(r.URL.Path, "/")
	if len(pieces) != 5 {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if pieces[1] != "update" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	mType := pieces[2]
	mName := pieces[3]
	mValue := pieces[4]

	if mName == "" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	switch mType {
	case "gauge":
		val, err := strconv.ParseFloat(mValue, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		h.storage.UpdateGauge(mName, val)

	case "counter":
		val, err := strconv.ParseInt(mValue, 10, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		h.storage.UpdateCounter(mName, val)

	default:
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
