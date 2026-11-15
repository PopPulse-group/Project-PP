// Package ai 定义 AI 菜谱推荐适配器。
// 约定：AI 只读商品数据；输出必须是可校验的结构化结果；
// 缺货、超预算或无法匹配必须显式标识，不允许编造商品、价格和库存；
// Mock 与真实调用的结果通过响应中的 mock 字段明确区分。
package ai

import (
	"context"
	"strconv"
	"strings"

	"vagetable-gozero/internal/types"
)

// SkuInfo AI 匹配用的在售商品只读视图（由商品模块提供）
type SkuInfo struct {
	Id    int64
	Name  string
	Unit  string
	Price int64 // 整数分
	Stock int64
}

// MatchInput 一次菜谱生成的输入
type MatchInput struct {
	PeopleCount int64
	BudgetCents int64
	Tastes      []string
	Avoids      []string
	Skus        []SkuInfo // 在售商品清单
}

// Provider AI 能力适配器接口；真实模型接入后实现同一接口并切换配置
type Provider interface {
	// GenerateMealPlan 返回结构化菜谱与 SKU 匹配结果（匹配由 Go 完成，保证不虚构价格库存）
	GenerateMealPlan(ctx context.Context, in MatchInput) (types.MealPlanResp, error)
}

// MockProvider 固定模板菜谱，不依赖外部服务；响应 mock=true
type MockProvider struct{}

func NewMockProvider() *MockProvider { return &MockProvider{} }

type mockTemplate struct {
	id, title, cuisine string
	ingredients        []string
	steps              []string
}

var mockTemplates = []mockTemplate{
	{"r-mock-001", "番茄炒蛋", "家常", []string{"番茄", "鸡蛋"}, []string{"番茄切块", "鸡蛋炒散", "合炒调味"}},
	{"r-mock-002", "青椒土豆丝", "家常", []string{"青椒", "土豆"}, []string{"土豆切丝泡水", "青椒切丝", "大火快炒"}},
	{"r-mock-003", "清蒸鲈鱼", "清淡", []string{"鲈鱼"}, []string{"鲈鱼处理干净", "大火蒸8分钟", "淋热油与生抽"}},
	{"r-mock-004", "香菇滑鸡", "家常", []string{"香菇", "鸡胸肉"}, []string{"鸡肉腌制", "香菇切片", "滑炒收汁"}},
}

// GenerateMealPlan 固定菜谱按名称关键词匹配在售 SKU；每 2 人一份用量；不可用时 available=false
func (p *MockProvider) GenerateMealPlan(ctx context.Context, in MatchInput) (types.MealPlanResp, error) {
	var avoidHint string
	skip := map[int]bool{}
	for _, a := range in.Avoids {
		if a == "" {
			continue
		}
		for i, t := range mockTemplates {
			if strings.Contains(t.title, a) || containsAny(t.ingredients, a) {
				skip[i] = true
				avoidHint += "已按忌口过滤：" + t.title + "（含" + a + "）；"
			}
		}
	}

	// 口味偏好命中的菜谱优先
	order := make([]int, 0, len(mockTemplates))
	for i := range mockTemplates {
		if !skip[i] {
			order = append(order, i)
		}
	}
	for _, taste := range in.Tastes {
		if taste == "" {
			continue
		}
		for pos, i := range order {
			t := mockTemplates[i]
			if strings.Contains(t.title, taste) || strings.Contains(t.cuisine, taste) {
				order[0], order[pos] = order[pos], order[0]
				break
			}
		}
	}
	if len(order) > 3 {
		order = order[:3]
	}

	var recipes []types.RecipeWithSkus
	for _, i := range order {
		t := mockTemplates[i]
		var items []types.SkuMatchItem
		sum := int64(0)
		for _, ing := range t.ingredients {
			qty := int64(1)
			if in.PeopleCount >= 4 {
				qty = 2
			}
			it := types.SkuMatchItem{Ingredient: ing, Quantity: qty, Available: false}
			if sku := findSku(in.Skus, ing); sku != nil {
				it.SkuId = strconv.FormatInt(sku.Id, 10)
				it.ProductName = sku.Name
				it.UnitPrice = sku.Price
				it.Amount = sku.Price * qty
				it.Available = sku.Stock >= qty
				sum += it.Amount
			}
			items = append(items, it)
		}
		recipes = append(recipes, types.RecipeWithSkus{
			Recipe: types.RecipeDish{
				RecipeId: t.id,
				Title:    t.title,
				Cuisine:  t.cuisine,
				Steps:    t.steps,
			},
			Items:        items,
			TotalAmount:  sum,
			WithinBudget: sum <= in.BudgetCents,
		})
	}

	return types.MealPlanResp{
		SessionId:   "mock-session",
		PeopleCount: in.PeopleCount,
		BudgetCents: in.BudgetCents,
		Recipes:     recipes,
		AvoidHint:   avoidHint,
		Mock:        true,
	}, nil
}

func findSku(skus []SkuInfo, keyword string) *SkuInfo {
	for i := range skus {
		if strings.Contains(skus[i].Name, keyword) {
			return &skus[i]
		}
	}
	return nil
}

func containsAny(list []string, sub string) bool {
	for _, s := range list {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
