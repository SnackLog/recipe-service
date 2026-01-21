package models

import "time"

type Recipe struct {
	Id                int                `db:"id" json:"-"`
	Name              string             `db:"name" json:"name"`
	Unit              string             `db:"unit" json:"unit"`
	Username          string             `db:"username" json:"-"`
	CreatedAt         time.Time          `db:"created_at" json:"created_at"`
	Ingredients       []Ingredient       `db:"-" json:"ingredients"`
	CustomIngredients []CustomIngredient `db:"-" json:"custom_ingredients"`
}
