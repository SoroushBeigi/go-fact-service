package main

import (
	"context"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	svc := NewFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc, logger)

	_, err := svc.GetFact(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

}
