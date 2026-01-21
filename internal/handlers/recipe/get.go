package recipe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RecipeController) Get(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
