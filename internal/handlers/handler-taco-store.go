package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"net/http"
)

// TacoStoreAPIHandler holds endpoints for taco list
type TacoStoreAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

// NewTacoStoreAPIHandler returns a new instance of TacoStoreAPIHandler.
func NewTacoStoreAPIHandler(mem *memstore.MemStore) *TacoStoreAPIHandler {
	h := &TacoStoreAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/api/taco-list", h.GetTacoList)

	return h
}

// GetTacoList gets taco list
func (h *TacoStoreAPIHandler) GetTacoList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	list := h.Store.SelectTacoList()
	json.NewEncoder(w).Encode(list)
}
