package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"net/http"
)

type TacoStoreAPIHandler struct {
	Store *memstore.TacoStore
	*httprouter.Router
}

func NewTacoStoreAPIHandler(mem *memstore.TacoStore) *TacoStoreAPIHandler {
	h := &TacoStoreAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/api/taco-list", h.GetTacoList)

	return h
}

func (h *TacoStoreAPIHandler) GetTacoList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	list := h.Store.SelectTacoList()
	json.NewEncoder(w).Encode(list)
}
