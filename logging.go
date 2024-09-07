package main

import (
	"context"
	"log"
	"time"
)

type LoggingService struct {
	next   Service
	logger *log.Logger
}

func NewLoggingService(next Service, logger *log.Logger) Service {
	return &LoggingService{next: next, logger: logger}
}

func (s *LoggingService) GetFact(ctx context.Context) (fact *Fact, err error) {
	defer func(start time.Time) {
		s.logger.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetFact(ctx)
}
