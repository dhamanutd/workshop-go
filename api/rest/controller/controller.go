package controller

import (
	"net/http"
)

type V1OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type ContextKey string
