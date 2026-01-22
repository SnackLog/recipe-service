package models

// Ingredient represents an ingredient in the DB and in requests.
type Ingredient struct {
	Id           int     `db:"id" json:"-"`
	IngredientId int     `db:"ingredient_id" json:"ingredient_id" binding:"required,min=1"`
	Quantity     float64 `db:"quantity" json:"quantity" binding:"required,gt=0"`
}

// CustomIngredient represents a custom ingredient in the DB and in requests, corresponding to a user-defined ingredient.
type CustomIngredient struct {
	Id                 int     `db:"id" json:"-"`
	CustomIngredientId int     `db:"custom_ingredient_id" json:"custom_ingredient_id" binding:"required,min=1"`
	Quantity           float64 `db:"quantity" json:"quantity" binding:"required,gt=0"`
}
