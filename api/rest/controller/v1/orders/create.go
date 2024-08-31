package orders_controller

import (
	"encoding/json"
	"net/http"

	"github.com/dhamanutd/golang-clean-architecture/internal/entities"
	"github.com/dhamanutd/golang-clean-architecture/pkg/transport"
)

func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		err error
		req entities.CreateOrderRequest
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		transport.SendResponseError(w, err)
		return
	}

	if err = c.validator.Struct(req); err != nil {
		transport.SendResponseError(w, err)
		return
	}

	if err = c.usecase.Create(ctx, &req); err != nil {
		transport.SendResponseError(w, err)
		return
	}

	transport.SendResponseOK(w, nil)
}
