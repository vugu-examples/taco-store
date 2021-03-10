package state

import (
	"errors"
	"fmt"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"golang.org/x/sync/singleflight"
	"log"
	"net/http"
)

// CardAPI holds sku status.
type CardAPI struct {
	Card []memstore.Taco
	g    singleflight.Group
}

// GetCard fetches one service agreement record
func (c *CardAPI) GetCard() ([]memstore.Taco, bool, error) {
	//use singleflight to deduplicate
	updated, err, _ := c.g.Do("/api/card", func() (interface{}, error) {
		// if we already have the data, don't pull again
		if c.Card != nil {
			return false, nil
		}
		url := "/api/card"
		res, err := Get(url)
		if err != nil {
			log.Printf("Error GetCard() %v", err)
			return false, err
		}
		if res.StatusCode != http.StatusOK {
			err = errors.New(fmt.Sprintf("Get %s returned status code %v", url, res.StatusCode))
			return false, err
		}
		err = Decoder(res, &c.Card)
		if err != nil {
			return false, err
		}
		return true, nil
	})
	return c.Card, updated.(bool), err
}

// CompleteAtmStateChangeApproval send a post request to create a new queue record
func (c *CardAPI) PostCard(payload memstore.Taco) error {

	url := "/api/card"
	res, err := Post(url, "application/json", payload)
	if err != nil {
		log.Printf("Error CompleteAtmStateChangeApproval: %v", err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Post %s returned status code %v", url, res.StatusCode))
		return err
	}
	return nil
}

func LoadCardAPI() *CardAPI {
	return &CardAPI{}
}

type CardAPISetter interface {
	CardAPISet(v *CardAPI)
}

type CardAPIRef struct {
	*CardAPI
}

func (r *CardAPIRef) CardAPISet(v *CardAPI) {
	r.CardAPI = v
}
