package opening

import (
	"net/http"
	"github.com/peruccii/gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request :=  handler.UpdateOpeningRequest{}
	
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.ErrF("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest, 
		handler.ErrParamisRequired("id", "QueryParameter").
		Error())
		return
	}
	opening := entity.Opening{}

	if err := handler.Db.First(&opening, id); err != nil {
		handler.SendError(ctx, http.StatusNotFound, handler.NotFoundId("opening", id))
		return 
	}
	// Update company
	if request.Role != "" {
			opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if err := handler.Db.Save(&opening).Error; err != nil {
		handler.Logger.ErrF("error updating opening %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	handler.SendSuccess(ctx, "updating-opening", opening)
}