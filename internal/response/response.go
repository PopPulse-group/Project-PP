// Package response 统一 HTTP 响应结构：{code, message, data, trace_id}。
package response

import (
	"context"
	"errors"
	"net/http"

	"poppulse/internal/xerr"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	TraceId string `json:"trace_id"`
}

func Success(data any, traceId string) Body {
	return Body{Code: 0, Message: "success", Data: data, TraceId: traceId}
}

// ErrorHandler 将错误映射为统一错误响应，供 httpx.SetErrorHandlerCtx 使用
func ErrorHandler(ctx context.Context, err error) (int, any) {
	var e *xerr.CodeError
	if errors.As(err, &e) {
		return e.Status, Body{Code: e.Code, Message: e.Msg, TraceId: traceIdFrom(ctx)}
	}
	return http.StatusInternalServerError, Body{Code: xerr.ErrInternal.Code, Message: xerr.ErrInternal.Msg, TraceId: traceIdFrom(ctx)}
}

// OkHandler 将 logic 返回值包装为统一成功响应，供 httpx.SetOkHandler 使用
func OkHandler(ctx context.Context, data any) any {
	return Success(data, traceIdFrom(ctx))
}

func traceIdFrom(ctx context.Context) string {
	if v, ok := ctx.Value("X-Trace-Id").(string); ok {
		return v
	}
	return ""
}
