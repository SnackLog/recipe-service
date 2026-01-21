package recipes

import (
	"database/sql"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

func GetById(db *sql.DB, id int) (*models.Recipe, error) {
	var recipe models.Recipe
	query := "SELECT id, name, unit, created_at, username FROM recipes WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&recipe.Id, &recipe.Name, &recipe.Unit, &recipe.CreatedAt, &recipe.Username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	ingredientQuery := "SELECT id, ingredient_id, quantity FROM ingredients WHERE recipe_id = $1"
	rows, err := db.Query(ingredientQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient models.Ingredient
		if err := rows.Scan(&ingredient.Id, &ingredient.IngredientId, &ingredient.Quantity); err != nil {
			return nil, err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	customIngredientQuery := "SELECT id, custom_ingredient_id, quantity FROM custom_ingredients WHERE recipe_id = $1"
	customRows, err := db.Query(customIngredientQuery, id)
	if err != nil {
		return nil, err
	}
	defer customRows.Close()

	for customRows.Next() {
		var customIngredient models.CustomIngredient
		if err := customRows.Scan(&customIngredient.Id, &customIngredient.CustomIngredientId, &customIngredient.Quantity); err != nil {
			return nil, err
		}
		recipe.CustomIngredients = append(recipe.CustomIngredients, customIngredient)
	}

	return &recipe, nil
}
