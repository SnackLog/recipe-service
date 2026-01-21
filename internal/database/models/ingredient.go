package models

type Ingredient struct {
	Id           int     `db:"id" json:"-"`
	IngredientId int     `db:"ingredient_id" json:"ingredient_id"`
	Quantity     float64 `db:"quantity" json:"quantity"`
}

type CustomIngredient struct {
	Id                 int     `db:"id" json:"-"`
	CustomIngredientId int     `db:"custom_ingredient_id" json:"custom_ingredient_id"`
	Quantity           float64 `db:"quantity" json:"quantity"`
}
