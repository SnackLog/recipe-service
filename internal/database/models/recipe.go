package models

import "time"

type Recipe struct {
	Id                int                `db:"id"`
	Name              string             `db:"name"`
	Unit              string             `db:"unit"`
	Username          string             `db:"username"`
	CreatedAt         time.Time          `db:"created_at"`
	Ingredients       []Ingredient       `db:"-"`
	CustomIngredients []CustomIngredient `db:"-"`
}
