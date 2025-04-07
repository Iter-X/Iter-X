package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func LanguageRepositoryImplToEntity(po *ent.Language) *do.Language {
	if po == nil {
		return nil
	}

	return &do.Language{
		Code:       po.Code,
		Name:       po.Name,
		NativeName: po.NativeName,
		Enabled:    po.Enabled,
		CreatedAt:  po.CreatedAt,
		UpdatedAt:  po.UpdatedAt,
	}
}
