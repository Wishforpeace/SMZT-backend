package jottings

import (
	"SMZT/handler"
	"SMZT/log"
	_ "SMZT/model/jottings"
	"SMZT/pkg/errno"
	service "SMZT/service/Jotting"
	"SMZT/util"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

// @Summary 获取随笔
// @Description 获取用户之前所发布的随笔
// @Tags jotting
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "token 用户令牌"
// @Success 200  "Success"
// @Failure 400 {string} json  "{"Code":400, "Message":"Error occurred while binding the request body to the struct","Data":nil}"
// @Failure 500 {string} json  "{"Code":500,"Message":"Database error","Data":nil}"
// @Router /jotting [get]
func GetJotting(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	StudentID := c.MustGet("student_id").(string)
	//StudentID := "2021214266"
	jotting, err := service.GetJotting(StudentID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, jotting)
}
