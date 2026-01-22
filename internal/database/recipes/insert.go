package recipes

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

// InsertWithTransactionAt Inserts recipe at id as part of transaction tx
func InsertWithTransactionAt(tx *sql.Tx, recipe *models.Recipe, id int) error {
	recipeId, err := insertRecipeAt(tx, recipe, id)
	if err != nil {
		return fmt.Errorf("Unable to insert recipe: %v", err)
	}

	err = insertIngredients(tx, recipeId, recipe.Ingredients)
	if err != nil {
		return fmt.Errorf("Unable to insert ingredients: %v", err)
	}

	err = insertCustomIngredients(tx, recipeId, recipe.CustomIngredients)
	if err != nil {
		return fmt.Errorf("Unable to insert custom ingredients: %v", err)
	}
	return nil
}

// InsertWithTransaction Inserts a new recipe using transaction tx, automatically determines the new ID of the recipe and returns it
func InsertWithTransaction(tx *sql.Tx, recipe *models.Recipe) (int, error) {
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
	return recipeId, nil
}

// Insert atomically inserts a new recipe and returns the new id
func Insert(db *sql.DB, recipe *models.Recipe) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return -1, fmt.Errorf("Unable to begin transaction: %v", err)
	}
	defer tx.Rollback()

	recipeId, err := InsertWithTransaction(tx, recipe)
	if err != nil {
		return -1, fmt.Errorf("Unable to insert recipe: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("Unable to commit transaction: %v", err)
	}

	return recipeId, nil
}

// insertRecipeAt utilizes tx to insert a recipe at id, will not handle ingredients
func insertRecipeAt(tx *sql.Tx, recipe *models.Recipe, id int) (int, error) {
	var recipeId int
	query := "INSERT INTO recipes (id, name, unit, username) VALUES ($1, $2, $3, $4) RETURNING id"
	err := tx.QueryRow(query, id, recipe.Name, recipe.Unit, recipe.Username).Scan(&recipeId)
	if err != nil {
		return -1, err
	}
	return recipeId, nil
}

// insertRecipe inserts a new recipe using tx, will not handle ingredients, returns the new recipe id
func insertRecipe(tx *sql.Tx, recipe *models.Recipe) (int, error) {
	var recipeId int
	query := "INSERT INTO recipes (name, unit, username) VALUES ($1, $2, $3) RETURNING id"
	err := tx.QueryRow(query, recipe.Name, recipe.Unit, recipe.Username).Scan(&recipeId)
	if err != nil {
		return -1, err
	}
	return recipeId, nil
}

// insertCustomIngredients inserts custom ingredients for recipeId using tx
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

// insertIngredients inserts ingredients for recipeId using tx
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
