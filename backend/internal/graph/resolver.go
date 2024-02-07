package graph

import (
	"github.com/tomhaerter/course-recommender/internal/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *database.Db
}
