package rest

import (
	"net/http"
	"time"

	"github.com/dhamanutd/golang-clean-architecture/api/rest/controller"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type RouterModule struct {
	V1OrderController controller.V1OrderController
}

func (module *RouterModule) Validate() error {
	return nil
}

func NewRouter(module RouterModule) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", module.V1OrderController.Create)
		})

	})

	return r
}
