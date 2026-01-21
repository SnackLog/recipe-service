package recipes

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

func Insert(db *sql.DB, recipe *models.Recipe) (int, error) {
	recipeId, err := insertRecipe(db, recipe)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert recipe: %v", err)
	}

	err = insertIngredients(db, recipeId, recipe.Ingredients)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert ingredients: %v", err)
	}

	err = insertCustomIngredients(db, recipeId, recipe.CustomIngredients)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert custom ingredients: %v", err)
	}

	return recipeId, nil
}

func insertRecipe(db *sql.DB, recipe *models.Recipe) (int, error) {
	var recipeId int
	query := "INSERT INTO recipes (name, unit, username) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(query, recipe.Name, recipe.Unit, recipe.Username).Scan(&recipeId)
	if err != nil {
		return -1, err
	}
	return recipeId, nil
}

func insertCustomIngredients(db *sql.DB, recipeId int, customIngredients []models.CustomIngredient) error {
	query := "INSERT INTO custom_ingredients (recipe_id, custom_ingredient_id, quantity) VALUES ($1, $2, $3)"
	for _, ingredient := range customIngredients {
		_, err := db.Exec(query, recipeId, ingredient.CustomIngredientId, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertIngredients(db *sql.DB, recipeId int, ingredients []models.Ingredient) error {
	query := "INSERT INTO ingredients (recipe_id, ingredient_id, quantity) VALUES ($1, $2, $3)"
	for _, ingredient := range ingredients {
		_, err := db.Exec(query, recipeId, ingredient.IngredientId, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
