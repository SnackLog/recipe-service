package recipe

import "database/sql"

// RecipeController Contains the database connection for recipe handlers
type RecipeController struct {
	DB *sql.DB
}
