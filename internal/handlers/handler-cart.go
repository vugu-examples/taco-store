package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"net/http"
)

type CardAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

func NewCardAPIHandler(mem *memstore.MemStore) *CardAPIHandler {
	h := &CardAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/api/card", h.GetCard)
	h.Router.POST("/api/card", h.PostCard)

	return h
}

func (h *CardAPIHandler) PostCard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var q memstore.Taco
	json.NewDecoder(r.Body).Decode(&q)
	h.Store.PostCard(q)
}
func (h *CardAPIHandler) GetCard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	card := h.Store.SelectCard()
	json.NewEncoder(w).Encode(card)
}
