package models

type Ingredient struct {
	Id           int     `db:"id" json:"-"`
	IngredientId int     `db:"ingredient_id" json:"ingredient_id" binding:"required,min=1"`
	Quantity     float64 `db:"quantity" json:"quantity" binding:"required,gt=0"`
}

type CustomIngredient struct {
	Id                 int     `db:"id" json:"-"`
	CustomIngredientId int     `db:"custom_ingredient_id" json:"custom_ingredient_id" binding:"required,min=1"`
	Quantity           float64 `db:"quantity" json:"quantity" binding:"required,gt=0"`
}
