package opening

import (
	"net/http"
	"github.com/peruccii/gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := handler.CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.ErrF("validantion error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return 
	}

	opening := entity.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Link: request.Link,
		Remote: *request.Remote,
		Salary: request.Salary,
	}

	handler.Logger.InfoF("request received %+v", request)

	if err := handler.Db.Create(&opening).Error; err != nil{
		handler.Logger.ErrF("error creating opening %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error creating opnening on database")
		return 
	}

	handler.SendSuccess(ctx, "create-opening", opening)
}