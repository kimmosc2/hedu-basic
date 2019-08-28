package controller

import (
	"github.com/gin-gonic/gin"
	"hedu-basic/serialize"
	"hedu-basic/service"
	"hedu-basic/util/ecode"
)

func ChangePassWord(c *gin.Context) {
	var pwdService service.ChangePasswordService
	if err := c.ShouldBind(&pwdService); err == nil {
		if err := pwdService.ChangePassword(); err == ecode.NoError {
			c.JSON(200, serialize.BuildAutoResponse(ecode.NoError, nil))
		} else {
			c.JSON(200, serialize.BuildErrorResponse(err))
		}
	} else {
		c.JSON(200, serialize.BuildErrorResponse(ecode.DataFormatError))
	}
}
