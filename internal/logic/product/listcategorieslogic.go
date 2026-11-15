// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package product

import (
	"context"

	"vagetable-gozero/internal/svc"
	"vagetable-gozero/internal/types"
	"vagetable-gozero/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分类列表
func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoriesLogic) ListCategories() (resp *types.CategoryListResp, err error) {
	return nil, xerr.ErrNotImplemented
}
