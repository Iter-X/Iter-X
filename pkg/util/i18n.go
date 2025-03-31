package util

import (
	"context"

	"github.com/iter-x/iter-x/internal/common/cnst"
)

// GetLocalizedName returns the localized name based on context language
// If context language is zh-CN, returns zhName, otherwise returns enName
func GetLocalizedName(ctx context.Context, enName, zhName string) string {
	if lang := ctx.Value(cnst.CtxKeyLang); lang == cnst.LangZhCn {
		return zhName
	}
	return enName
}
