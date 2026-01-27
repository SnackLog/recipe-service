package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/SnackLog/recipe-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Delete handles DELETE /recipe/:id requests to delete a recipe by ID
// @Summary Delete a recipe
// @Description Deletes a recipe by ID for the authenticated user.
// @Tags recipe
// @Produce json
// @Param id path int true "Recipe ID"
// @Success 204 "No Content"
// @Failure 400 {object} handlers.Error
// @Failure 401 {object} handlers.Error
// @Failure 404 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Security ApiKeyAuth
// @Router /recipe/{id} [delete]
func (rc *RecipeController) Delete(c *gin.Context) {
	username := c.GetString("username")
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting recipe ID: %v", err)
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "invalid recipe ID"})
		return
	}

	result, err := recipes.DeleteRecipe(rc.DB, recipeID, username)
	if err != nil {
		log.Printf("Error deleting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "failed to delete recipe"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "failed to delete recipe"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, handlers.Error{Error: "recipe not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
