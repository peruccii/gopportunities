package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data": data,
	})
}


func NotFoundId(value string, p any) string{
	not_found_msg := fmt.Sprintf("%s with id: %s not found", value, p)
	return not_found_msg
}