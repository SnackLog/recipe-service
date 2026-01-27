package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/SnackLog/recipe-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// GetID handles GET /recipe/:id requests to retrieve a recipe by ID
// @Summary Get a recipe by ID
// @Description Retrieves a recipe by ID for the authenticated user.
// @Tags recipe
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 200 {object} models.Recipe
// @Failure 400 {object} handlers.Error
// @Failure 401 {object} handlers.Error
// @Failure 404 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Security ApiKeyAuth
// @Router /recipe/{id} [get]
func (rc *RecipeController) GetID(c *gin.Context) {
	idParam := c.Param("id")
	username := c.GetString("username")
	recipeId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid recipe ID"})
		return
	}

	recipe, err := recipes.GetById(rc.DB, recipeId)

	if err != nil {
		log.Println("Error retrieving recipe:", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to retrieve recipe"})
		return
	}
	if recipe == nil || recipe.Username != username {
		c.JSON(http.StatusNotFound, handlers.Error{Error: "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}
