package jottings

import (
	"SMZT/handler"
	"SMZT/log"
	"SMZT/pkg/errno"
	"SMZT/service/Jotting"
	"SMZT/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 发布随笔
// @Description 用户根据每日感受写随想笔记
// @Tags jotting
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "token 用户令牌"
// @Success 200  "Success"
// @Failure 400 {string} json  "{"Code":400, "Message":"Error occurred while binding the request body to the struct","Data":nil}"
// @Failure 500 {string} json  "{"Code":500,"Message":"Database error","Data":nil}"
// @Router /jotting [post]
func PostJotting(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	StudentID := c.MustGet("student_id").(string)
	//StudentID := "2021214266"

	var req JottingRequest
	if err := c.ShouldBind(&req); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error(), handler.GetLine())
		return
	}

	if err := Jotting.PostJotting(StudentID, req.Title, req.Content); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "Success")

}
