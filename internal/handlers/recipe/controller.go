package recipe

import "database/sql"

type RecipeController struct {
	DB *sql.DB
}
