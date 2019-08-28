package ecode

const (
	// 全局错误码
	NoError = 0 // 操作成功

	// 官方服务错误码
	DataFormatError      = 40001 // 数据格式错误(如用户名或密码为空)
	UserNameEmpty        = 40002 // 用户名为空
	PasswordEmpty        = 40003 // 密码为空
	ChangePasswordFailed = 40004 // 修改密码失败(一般是账号不存在导致的)

	// 服务内部错误
	OfficialServiceUnavailability = 50001 // 请求失败,服务不可用,请检查是不是官网挂了
)

var Ecode = map[int]string{
	// 全局错误码
	NoError: "操作成功",
	// 官方服务错误码
	DataFormatError:      "数据格式错误", // 数据格式错误(如用户名或密码为空)
	UserNameEmpty:        "用户名为空",
	PasswordEmpty:        "密码为空",
	ChangePasswordFailed: "修改密码失败",

	// 服务内部错误
	OfficialServiceUnavailability: "服务内部错误,请联系管理员",
}
