package recipes

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

func Insert(db *sql.DB, recipe *models.Recipe) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return -1, fmt.Errorf("Unable to begin transaction: %v", err)
	}
	defer tx.Rollback()

	recipeId, err := insertRecipe(tx, recipe)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert recipe: %v", err)
	}

	err = insertIngredients(tx, recipeId, recipe.Ingredients)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert ingredients: %v", err)
	}

	err = insertCustomIngredients(tx, recipeId, recipe.CustomIngredients)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert custom ingredients: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("Unable to commit transaction: %v", err)
	}

	return recipeId, nil
}

func insertRecipe(tx *sql.Tx, recipe *models.Recipe) (int, error) {
	var recipeId int
	query := "INSERT INTO recipes (name, unit, username) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRow(query, recipe.Name, recipe.Unit, recipe.Username).Scan(&recipeId)
	if err != nil {
		return -1, err
	}
	return recipeId, nil
}

func insertCustomIngredients(tx *sql.Tx, recipeId int, customIngredients []models.CustomIngredient) error {
	query := "INSERT INTO custom_ingredients (recipe_id, custom_ingredient_id, quantity) VALUES ($1, $2, $3)"
	for _, ingredient := range customIngredients {
		_, err := tx.Exec(query, recipeId, ingredient.CustomIngredientId, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertIngredients(tx *sql.Tx, recipeId int, ingredients []models.Ingredient) error {
	query := "INSERT INTO ingredients (recipe_id, ingredient_id, quantity) VALUES ($1, $2, $3)"
	for _, ingredient := range ingredients {
		_, err := tx.Exec(query, recipeId, ingredient.IngredientId, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
