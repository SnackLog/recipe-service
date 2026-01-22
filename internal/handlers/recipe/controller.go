package recipe

import "database/sql"

// RecipeController is a controller that holds the database connection for recipe handlers.
type RecipeController struct {
	DB *sql.DB
}
