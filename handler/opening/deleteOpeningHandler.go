package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
	"github.com/peruccii/gopportunities/handler"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, 
		handler.ErrParamisRequired("id", "queryParameter").
		Error())
		return
	}
	opening := entity.Opening{}
	// Find Opening
	if err := handler.Db.First(&opening, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound, handler.NotFoundId("opening", id))
		return
	}
	// Delete Opening
	if err := handler.Db.Delete(&opening).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, handler.NotFoundId("opening deleting", id))
		return
	}
	handler.SendSuccess(ctx, "delete-opening", opening)
}