package http

import (
	"net/http"
	"sync"
)

type ServeMux struct {
	// なぜ必要なの？
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest
	hosts bool       // whether any patterns contain hostname
}

type muxEntry struct {
	h       http.Handler
	pattern string
}
