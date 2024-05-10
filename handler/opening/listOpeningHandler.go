package opening

import (
	"net/http"
	"github.com/peruccii/gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []entity.Opening{}

	 if err := handler.Db.Find(&openings).Error; err != nil {
			handler.SendError(ctx, http.StatusInternalServerError, "opening does not exists")
			return
		}
		handler.SendSuccess(ctx, "list-openings", openings)
}