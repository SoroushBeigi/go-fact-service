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
		s.logger.Printf("fact=%v err=%v took=%v", fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetFact(ctx)
}
