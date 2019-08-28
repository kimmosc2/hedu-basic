package service

import (
	"hedu-basic/util/component/official/pwd"
	"hedu-basic/util/ecode"
)

type ChangePasswordService struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (cs *ChangePasswordService) ChangePassword() int {
	if err := cs.checkAccount();err!= ecode.NoError {
		return err
	}
	return pwd.ChangePassWord(cs.Username,cs.Password)
}

// 验证账密码是否为空
func (cs *ChangePasswordService) checkAccount() int {
	if cs.Username == "" {
		return ecode.UserNameEmpty
	} else if cs.Password == "" {
		return ecode.PasswordEmpty
	}
	return ecode.NoError
}
