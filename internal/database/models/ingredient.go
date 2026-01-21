package models

type Ingredient struct {
	Id           int     `db:"id"`
	IngredientId int     `db:"ingredient_id"`
	Quantity     float64 `db:"quantity"`
}

type CustomIngredient struct {
	Id                 int    `db:"id"`
	CustomIngredientId int    `db:"custom_ingredient_id"`
	Quantity     float64 `db:"quantity"`
}
