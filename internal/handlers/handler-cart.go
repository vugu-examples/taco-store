package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"net/http"
)

type CartAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

func NewCartAPIHandler(mem *memstore.MemStore) *CartAPIHandler {
	h := &CartAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/api/cart", h.GetCart)
	h.Router.POST("/api/cart", h.PostCart)

	return h
}

func (h *CartAPIHandler) PostCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var q memstore.Taco
	json.NewDecoder(r.Body).Decode(&q)
	h.Store.PostCart(q)
}
func (h *CartAPIHandler) GetCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Cart := h.Store.SelectCart()
	json.NewEncoder(w).Encode(Cart)
}
