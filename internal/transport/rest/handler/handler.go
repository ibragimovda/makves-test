package handler

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"makves/internal/entity"
	"net/http"
)

type Response struct {
	Success int    `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Service interface {
	GetItems(ids []int) ([]entity.Item, error)
}

type Handler struct {
	service Service
	log     *slog.Logger
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-items", h.GetItems)
	return mux
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data := r.URL.Query().Get("id")
	if data == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Error: "Не найден параметр запроса"})
		return
	}

	var ids []int
	err := json.Unmarshal([]byte(data), &ids)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Error: "Неверный параметр"})
		return
	}

	items, err := h.service.GetItems(ids)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(Response{Success: 0, Error: "Ошибка при получении данных"})
		h.log.Error(err.Error())
		return
	}
	json.NewEncoder(w).Encode(Response{Success: 1, Data: items})
}

func NewHandler(service Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}
