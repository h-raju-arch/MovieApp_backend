package httptransport

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
)

func (h Movie_handler) GetMovies(c *gin.Context) {

	ctx := c.Request.Context()
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie id needed"})
		return
	}

	_, err := uuid.FromString(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
	}

	lang := c.DefaultQuery("lang", "en")
	appendtoresponse := c.Query("append_to_response")
	var appends []string
	if appendtoresponse != "" {
		for _, s := range strings.Split(appendtoresponse, ",") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			appends = append(appends, s)
		}
	}

	res, err := h.svc.GetMovieById(ctx, id, lang, appends)

	if err != nil {
		fmt.Println("GetMovies error: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
