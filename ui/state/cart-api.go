package state

import (
	"errors"
	"fmt"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"github.com/vugu-examples/taco-store/ui/format"
	"golang.org/x/sync/singleflight"
	"log"
	"net/http"
)

// CartAPI holds API calls for Cart items
type CartAPI struct {
	Cart []memstore.Taco
	g    singleflight.Group
}

// GetCart fetches cart items and  uses singleflight to avoid multiple requests.
// please check that top-nav.vugu and cart.vugu call GetCart and ask for the same data in their initialization.
// see: https://pkg.go.dev/golang.org/x/sync/singleflight
func (c *CartAPI) GetCart() ([]memstore.Taco, bool, error) {
	//use singleflight to deduplicate
	updated, err, _ := c.g.Do("/api/Cart", func() (interface{}, error) {
		// if we already have the data, don't pull again
		if c.Cart != nil {
			return false, nil
		}
		url := "/api/cart"
		res, err := Get(url)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error GetCart() %v", err))
			return false, err
		}
		if res.StatusCode != http.StatusOK {
			err = errors.New(fmt.Sprintf("Get %s returned status code %v", url, res.StatusCode))
			return false, err
		}
		err = Decoder(res, &c.Cart)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error GetCart() Decoder  %v", err))
			return false, err
		}
		return true, nil
	})
	return c.Cart, updated.(bool), err
}

// PostCartItem creates a cart item
func (c *CartAPI) PostCartItem(payload memstore.Taco) error {

	url := "/api/cart"
	res, err := Post(url, "application/json", payload)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error PostCartItem() %v", err))
		return err
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("PostCartItem %s returned status code %v", url, res.StatusCode))
		return err
	}
	return nil
}

// DeleteCartItem deletes cart item
func (c *CartAPI) DeleteCartItem(payload []memstore.Taco) error {

	url := "/api/cart"
	res, err := Patch(url, payload)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error DeleteCartItem() %v", err))
		return err
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("DeleteCartItem %s returned status code %v", url, res.StatusCode))
		log.Printf("Error DeleteCartItem: %v", err)
		return err
	}
	return nil
}

// GetCartTotal calculates sum of cart items
func (c *CartAPI) GetCartTotal() string {
	if c.Cart == nil {
		return ""
	}
	var sum float32
	for _, item := range c.Cart {
		sum += item.Price
	}
	return format.Currency(sum)
}

// LoadCartAPI returns a new instance of CartAPI
func LoadCartAPI() *CartAPI {
	return &CartAPI{}
}

// CartAPISetter interface for wiring
type CartAPISetter interface {
	CartAPISet(v *CartAPI)
}

// CartAPIRef ref for wiring
type CartAPIRef struct {
	*CartAPI
}

// CartAPISet setter for wiring
func (r *CartAPIRef) CartAPISet(v *CartAPI) {
	r.CartAPI = v
}
