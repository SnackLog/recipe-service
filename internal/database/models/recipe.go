package models

import "time"

// Recipe represents a recipe created by the user, which may contain multiple Ingredients or CustomIngredients, as well as some metadata
type Recipe struct {
	Id                int                `db:"id" json:"-"`
	Name              string             `db:"name" json:"name" binding:"required,min=1,max=100"`
	Unit              string             `db:"unit" json:"unit" binding:"required,min=1,max=50"`
	Username          string             `db:"username" json:"-"`
	CreatedAt         time.Time          `db:"created_at" json:"created_at"`
	Ingredients       []Ingredient       `db:"-" json:"ingredients"`
	CustomIngredients []CustomIngredient `db:"-" json:"custom_ingredients"`
}
