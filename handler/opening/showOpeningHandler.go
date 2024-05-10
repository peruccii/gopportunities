package opening

import (
	"fmt"
	"net/http"
	"github.com/peruccii/gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		handler.SendError(ctx, http.StatusNotFound, 
		handler.ErrParamisRequired("id", "QueryParameter").
		Error())
		return
	}

	opening := entity.Opening{}

	if err := handler.Db.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening id: %s does not exists", id))
		return
	}

	handler.SendSuccess(ctx, "list-specific-opening", opening)
}