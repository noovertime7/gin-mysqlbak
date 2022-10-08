package globalError

type GlobalError struct {
	Code             int    `json:"code"`    // 业务码
	Message          string `json:"message"` // 业务码
	RealErrorMessage string `json:"err_msg"`
}

func (e *GlobalError) Error() string {
	return e.Message
}

// 2、定义errorCode
const (
	ServerError        = 10101 // Internal Server Error
	ParamBindError     = 10102 // 参数信息有误
	AuthorizationError = 10103 // 签名信息有误
	CallHTTPError      = 10104 // 调用第三方HTTP接口失败
	ResubmitMsg        = 10105 // 请勿重复提交

	HostAddError    = 20101 // 请勿重复提交
	HostDeleteError = 20102 // 请勿重复提交
	HostUpdateError = 20103 // 请勿重复提交
	HostGetError    = 20104 // 请勿重复提交
	HostCheckError  = 20105

	TaskAddError         = 20201 // 请勿重复提交
	TaskDeleteError      = 20202 // 请勿重复提交
	TaskUpdateError      = 20203 // 请勿重复提交
	TaskGetError         = 20204 // 请勿重复提交
	TaskNodeFound        = 20205
	TaskOverViewGetError = 20206
	TaskRestoreError     = 20207

	HistoryAddError    = 20301 // 请勿重复提交
	HistoryDeleteError = 20302 // 请勿重复提交
	HistoryUpdateError = 20303 // 请勿重复提交
	HistoryGetError    = 20304 // 请勿重复提交

	AdminCreateError             = 20401
	AdminListError               = 20402
	AdminDeleteError             = 20403
	AdminUpdateError             = 20404
	AdminResetPasswordError      = 20405
	AdminLoginError              = 20406
	AdminLogOutError             = 20407
	AdminModifyPasswordError     = 20408
	AdminModifyPersonalInfoError = 20409

	AgentRegisterError   = 20501
	AgentDeRegisterError = 20502
	AgentGetError        = 20503
	AgentGetAddressError = 20504

	BakStartError    = 20601 // 请勿重复提交
	BakStopError     = 20602 // 请勿重复提交
	BakStartAllError = 20603 // 请勿重复提交
	BakStopAllError  = 20604 // 请勿重复提交
)

// 3、定义errorCode对应的文本信息
var codeTag = map[int]string{
	ServerError:                  "Internal Server Error",
	ParamBindError:               "参数信息有误",
	AuthorizationError:           "签名信息有误",
	CallHTTPError:                "调用第三方 HTTP 接口失败",
	ResubmitMsg:                  "请勿重复提交",
	HostAddError:                 "主机添加失败，请联系管理员",
	HostDeleteError:              "主机删除失败，请联系管理员",
	HostUpdateError:              "主机更新失败，请联系管理员",
	HostGetError:                 "主机查询失败，请联系管理员",
	TaskAddError:                 "任务添加失败，请联系管理员",
	TaskDeleteError:              "任务删除失败，请联系管理员",
	TaskUpdateError:              "任务更新失败，请联系管理员",
	TaskGetError:                 "任务查询失败，请联系管理员",
	TaskNodeFound:                "备份任务为空",
	TaskOverViewGetError:         "获取任务总览失败",
	TaskRestoreError:             "还原任务失败",
	HistoryAddError:              "历史记录添加失败，请联系管理员",
	HistoryDeleteError:           "历史记录删除失败，请联系管理员",
	HistoryUpdateError:           "历史记录更新失败，请联系管理员",
	HistoryGetError:              "历史记录查询失败，请联系管理员",
	AdminCreateError:             "创建管理员失败",
	AdminListError:               "获取管理员列表页失败",
	AdminDeleteError:             "删除管理员失败",
	AdminUpdateError:             "更新管理员失败",
	AdminResetPasswordError:      "重置密码失败",
	AdminLoginError:              "登录失败",
	AdminLogOutError:             "退出失败",
	AdminModifyPasswordError:     "修改密码失败",
	AdminModifyPersonalInfoError: "修改个人信息失败",

	AgentRegisterError:   "客户端注册失败",
	AgentDeRegisterError: "客户端注销失败",
	AgentGetError:        "客户端查询失败，请联系管理员",
	AgentGetAddressError: "客户端获取地址失败",

	BakStartError:    "启动备份任务失败",
	BakStopError:     "停止备份任务失败",
	BakStartAllError: "批量启动备份任务失败",
	BakStopAllError:  "批量停止备份任务失败",
}

// NewGlobalError 4、新建自定义error实例化
func NewGlobalError(code int, err error) error {
	// 初次调用得用Wrap方法，进行实例化
	return &GlobalError{
		Code:             code,
		Message:          codeTag[code],
		RealErrorMessage: err.Error(),
	}
}
