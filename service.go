package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFact(ctx context.Context) (*Fact, error)
}

type FactService struct {
	url string
}

func NewFactService(url string) Service {
	return &FactService{url: url}
}

func (s *FactService) GetFact(ctx context.Context) (*Fact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fact := &Fact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}
	return fact, nil

}
