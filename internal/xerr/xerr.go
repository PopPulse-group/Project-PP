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
	ErrIdempotency    = New(40901, "幂等键重复但请求内容不同", 409)
	ErrUnauthorized   = New(40101, "未登录或 token 失效", 401)
	ErrForbidden      = New(40301, "无权限", 403)
	ErrNotFound       = New(40401, "资源不存在", 404)
	ErrNotImplemented = New(50101, "接口尚未实现（占位）", 501)
	ErrInternal       = New(50001, "服务内部错误", 500)
	ErrServiceDegrade = New(50301, "外部服务降级，稍后重试", 503)
	ErrRateLimited    = New(42901, "请求过于频繁", 429)
)

// 业务错误（按需使用 NewBiz 创建，或直接使用预定义）
func NewBiz(code int, msg string) *CodeError { return New(code, msg, 422) }

var (
	ErrStockInsufficient = New(40902, "库存不足", 409)
	ErrOrderStateInvalid = New(40903, "订单状态不允许此操作", 409)
	ErrMockDisabled      = NewBiz(40904, "模拟能力未启用")
	ErrAiOutputInvalid   = NewBiz(42210, "AI 输出校验失败")
	ErrIngredientMissing = NewBiz(42211, "部分食材无在售商品匹配")
	ErrLoginFailed       = New(40102, "用户名或密码错误", 401)
)
