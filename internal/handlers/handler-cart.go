package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"net/http"
)

// CartAPIHandler holds endpoints for cart items
type CartAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

// NewCartAPIHandler returns a new instance of CartAPIHandler.
func NewCartAPIHandler(mem *memstore.MemStore) *CartAPIHandler {
	h := &CartAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/api/cart", h.GetCart)
	h.Router.POST("/api/cart", h.PostCartItem)
	h.Router.PATCH("/api/cart", h.PatchCart)

	return h
}

// PostCartItem creates a cart item
func (h *CartAPIHandler) PostCartItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var q memstore.Taco
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		return
	}
	h.Store.PostCartItem(q)
}

// GetCart gets cart item list
func (h *CartAPIHandler) GetCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Cart := h.Store.SelectCart()
	err := json.NewEncoder(w).Encode(Cart)
	if err != nil {
		return
	}
}

// PatchCart updates cart items
func (h *CartAPIHandler) PatchCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var q []memstore.Taco
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		return
	}
	h.Store.PatchCart(q)
}
