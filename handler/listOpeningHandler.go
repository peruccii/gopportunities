package handler

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET OPENING", // Show in the body
	})
}