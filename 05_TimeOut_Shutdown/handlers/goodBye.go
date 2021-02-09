package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// GoodBye is a simple handler
type GoodBye struct {
	l *log.Logger
}

// NewGoodBye creates a new hello handler with the given logger
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle GoodBye requests")
	fmt.Fprintf(w, "GoodBye")
}
