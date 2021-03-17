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

type CartAPI struct {
	Cart []memstore.Taco
	g    singleflight.Group
}

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

func (c *CartAPI) PostCart(payload memstore.Taco) error {

	url := "/api/cart"
	res, err := Post(url, "application/json", payload)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error PostCart() %v", err))
		return err
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("PostCart %s returned status code %v", url, res.StatusCode))
		return err
	}
	return nil
}

func (c *CartAPI) PatchCart(payload []memstore.Taco) error {

	url := "/api/cart"
	res, err := Patch(url, payload)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error PatchCart() %v", err))
		return err
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("PatchCart %s returned status code %v", url, res.StatusCode))
		log.Printf("Error PatchCart: %v", err)
		return err
	}
	return nil
}

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

func LoadCartAPI() *CartAPI {
	return &CartAPI{}
}

type CartAPISetter interface {
	CartAPISet(v *CartAPI)
}

type CartAPIRef struct {
	*CartAPI
}

func (r *CartAPIRef) CartAPISet(v *CartAPI) {
	r.CartAPI = v
}
