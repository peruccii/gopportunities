package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peruccii/gopportunities/entity"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("validantion error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
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

	logger.InfoF("request received %+v", request)

	if err := db.Create(&opening).Error; err != nil{
		logger.ErrF("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opnening on database")
		return 
	}

	sendSuccess(ctx, "create-opening", opening)
}