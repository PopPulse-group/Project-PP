// Package xerr 定义统一业务错误码与错误类型。
//
// 约定（与接口文档一致）：
//   - 0            成功
//   - 400xx        参数错误（HTTP 400）
//   - 401xx        未认证（HTTP 401）
//   - 403xx        无权限（HTTP 403）
//   - 404xx        资源不存在（HTTP 404）
//   - 409xx        冲突/幂等冲突（HTTP 409）
//   - 300xx/422xx  业务规则拒绝（HTTP 422）
//   - 501xx        未实现（HTTP 501，D1 占位）
//   - 500xx        服务内部错误（HTTP 500）
//   - 900xx        依赖不可用（HTTP 503）
package xerr

import "fmt"

type CodeError struct {
	Code   int
	Msg    string
	Status int // HTTP 状态码
}

func (e *CodeError) Error() string { return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg) }

func New(code int, msg string, status int) *CodeError {
	return &CodeError{Code: code, Msg: msg, Status: status}
}

// 参数与权限
var (
	ErrInvalidParam   = New(40001, "参数错误", 400)
	ErrIdempotency    = New(40002, "幂等键冲突：相同幂等键但请求参数不同", 409)
	ErrUnauthorized   = New(40101, "未登录或登录已过期", 401)
	ErrForbidden      = New(40301, "无权限执行该操作", 403)
	ErrNotFound       = New(40401, "资源不存在", 404)
	ErrStateConflict  = New(40901, "状态冲突：当前状态不允许该操作", 409)
	ErrNotImplemented = New(50101, "接口尚未实现（D1 占位）", 501)
	ErrInternal       = New(50001, "服务内部错误", 500)
	ErrServiceDegrade = New(90002, "依赖服务不可用", 503)
)

// 业务错误（按需使用 NewBiz 创建，或直接使用预定义）
func NewBiz(code int, msg string) *CodeError { return New(code, msg, 422) }

var (
	ErrStockInsufficient   = NewBiz(30001, "库存不足")
	ErrOrderStateInvalid   = NewBiz(30002, "订单状态不允许该操作")
	ErrMockDisabled        = NewBiz(30003, "模拟能力未启用")
	ErrPaymentStateInvalid = NewBiz(30004, "支付状态不允许该操作")
	ErrAfterSaleInvalid    = NewBiz(30005, "售后状态不允许该操作")
	ErrRefundExceed        = NewBiz(30006, "退款金额超过可退额度")
	ErrUserRole            = NewBiz(30007, "账号角色不满足要求")
	ErrLoginFailed         = New(40003, "用户名或密码错误", 401)
)
