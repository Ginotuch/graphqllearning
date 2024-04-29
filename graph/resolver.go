package graph

import (
	"graphqllearning/graph/storage"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	database storage.Database
	mu       sync.RWMutex
}

func NewResolver(database storage.Database) *Resolver {
	return &Resolver{database: database}
}
