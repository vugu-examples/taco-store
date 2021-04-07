package state

import (
	"errors"
	"fmt"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"golang.org/x/sync/singleflight"
	"log"
	"net/http"
)

// TacoListAPI holds API calls for Taco List
type TacoListAPI struct {
	TacoList []memstore.Taco
	g        singleflight.Group
}

// GetTacoList fetches taco list
func (c *TacoListAPI) GetTacoList() ([]memstore.Taco, bool, error) {
	//use singleflight to deduplicate
	updated, err, _ := c.g.Do("/api/taco-list", func() (interface{}, error) {
		// if we already have the data, don't pull again
		if c.TacoList != nil {
			return false, nil
		}
		url := "/api/taco-list"
		res, err := Get(url)
		if err != nil {
			log.Printf("Error GetTacoList() %v", err)
			return false, err
		}
		if res.StatusCode != http.StatusOK {
			err = errors.New(fmt.Sprintf("Get %s returned status code %v", url, res.StatusCode))
			return false, err
		}
		err = Decoder(res, &c.TacoList)
		if err != nil {
			return false, err
		}
		return true, nil
	})
	return c.TacoList, updated.(bool), err
}

// LoadTacoListAPI returns a new instance of TacoListAPI
func LoadTacoListAPI() *TacoListAPI {
	return &TacoListAPI{}
}

// TacoListAPISetter interface for wiring
type TacoListAPISetter interface {
	TacoListAPISet(v *TacoListAPI)
}

// TacoListAPIRef ref for wiring
type TacoListAPIRef struct {
	*TacoListAPI
}

// TacoListAPISet setter for wiring
func (r *TacoListAPIRef) TacoListAPISet(v *TacoListAPI) {
	r.TacoListAPI = v
}
