package handler

import (
	d "test/golang/domain"
	r "test/golang/repository"
)

type Handler struct {
	Rp      r.Repo
	Service d.Services
}

func NewHandler(r r.Repo, s d.Services) *Handler {
	return &Handler{
		Rp:      r,
		Service: s,
	}
}
