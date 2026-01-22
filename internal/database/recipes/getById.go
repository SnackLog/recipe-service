package recipes

import (
	"database/sql"
	"fmt"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

// GetById returns a recipe, including all ingredients with id
func GetById(db *sql.DB, id int) (*models.Recipe, error) {
	recipe, err := getRecipe(db, id)
	if err != nil {
		return nil, fmt.Errorf("error getting recipe: %v", err)
	}
	if recipe == nil {
		return nil, nil
	}

	err = populateIngredients(db, id, recipe)
	if err != nil {
		return nil, fmt.Errorf("error getting ingredients: %v", err)
	}

	err = populateCustomIngredients(db, id, recipe)
	if err != nil {
		return nil, fmt.Errorf("error getting custom ingredients: %v", err)
	}

	return recipe, nil
}

// getRecipe aquires a recipe object from db using id
func getRecipe(db *sql.DB, id int) (*models.Recipe, error) {
	var recipe models.Recipe
	query := "SELECT id, name, unit, created_at, username FROM recipes WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&recipe.Id, &recipe.Name, &recipe.Unit, &recipe.CreatedAt, &recipe.Username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying recipe: %v", err)
	}
	return &recipe, nil
}

// populateIngredients populates the ingredients list on recipe
func populateIngredients(db *sql.DB, id int, recipe *models.Recipe) error {
	ingredientQuery := "SELECT id, ingredient_id, quantity FROM ingredients WHERE recipe_id = $1"
	rows, err := db.Query(ingredientQuery, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient models.Ingredient
		if err := rows.Scan(&ingredient.Id, &ingredient.IngredientId, &ingredient.Quantity); err != nil {
			return err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// populateCustomIngredients Populates CustomIngredients on recipe
func populateCustomIngredients(db *sql.DB, id int, recipe *models.Recipe) error {
	customIngredientQuery := "SELECT id, custom_ingredient_id, quantity FROM custom_ingredients WHERE recipe_id = $1"
	customRows, err := db.Query(customIngredientQuery, id)
	if err != nil {
		return err
	}
	defer customRows.Close()

	for customRows.Next() {
		var customIngredient models.CustomIngredient
		if err := customRows.Scan(&customIngredient.Id, &customIngredient.CustomIngredientId, &customIngredient.Quantity); err != nil {
			return err
		}
		recipe.CustomIngredients = append(recipe.CustomIngredients, customIngredient)
	}
	err = customRows.Err()
	if err != nil {
		return err
	}
	return nil
}
