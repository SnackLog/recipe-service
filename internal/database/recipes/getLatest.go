package recipes

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

func GetLatest(db *sql.DB, username string, limit int) ([]models.Recipe, error) {
	query := "SELECT id, name, unit, created_at, username FROM recipes WHERE username = $1 ORDER BY created_at DESC LIMIT $2"
	rows, err := db.Query(query, username, limit)
	if err != nil {
		return nil, fmt.Errorf("error querying latest recipes: %v", err)
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Unit, &recipe.CreatedAt, &recipe.Username); err != nil {
			return nil, fmt.Errorf("error scanning recipe: %v", err)
		}
		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over recipe rows: %v", err)
	}

	for i := range recipes {
		err = populateIngredients(db, recipes[i].Id, &recipes[i])
		if err != nil {
			return nil, fmt.Errorf("error populating ingredients for recipe %d: %v", recipes[i].Id, err)
		}

		err = populateCustomIngredients(db, recipes[i].Id, &recipes[i])
		if err != nil {
			return nil, fmt.Errorf("error populating custom ingredients for recipe %d: %v", recipes[i].Id, err)
		}
	}

	return recipes, nil
}
