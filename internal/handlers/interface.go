package handlers

import "context"

type Handler struct {
}

func NewHandler(ctx context.Context) *Handler {

	return &Handler{}
}
