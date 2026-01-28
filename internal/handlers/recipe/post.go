package recipe

import (
	"log"
	"net/http"

	"github.com/SnackLog/recipe-service/internal/database/models"
	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/SnackLog/recipe-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type recipePostResponse struct {
	RecipeID int `json:"recipe_id"`
}

// Post handles POST /recipe requests to create a new recipe
// @Summary Create a recipe
// @Description Creates a new recipe for the authenticated user.
// @Tags recipe
// @Accept json
// @Produce json
// @Param recipe body models.Recipe true "Recipe object"
// @Success 201 {object} map[string]int
// @Failure 400 {object} handlers.Error
// @Failure 401 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Security ApiKeyAuth
// @Router /recipe [post]
func (rc *RecipeController) Post(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request payload"})
		return
	}
	recipe.Username = c.GetString("username")

	recipeId, err := recipes.Insert(rc.DB, &recipe)
	if err != nil {
		log.Printf("Error inserting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, recipePostResponse{RecipeID: recipeId})

}
