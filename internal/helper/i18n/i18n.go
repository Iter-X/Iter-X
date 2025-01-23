package i18n

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/google/wire"
	"github.com/iter-x/iter-x/internal/conf"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path"
)

var ProviderSet = wire.NewSet(New)

func New(c *conf.I18N) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	suffix := ".toml"
	switch c.Format {
	case conf.I18NFormat_JSON:
		suffix = ".json"
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	case conf.I18NFormat_TOML:
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	default:
		panic("unknown i18n format")
	}
	files, err := os.ReadDir(c.Dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if path.Ext(file.Name()) != suffix {
			continue
		}
		bundle.MustLoadMessageFile(path.Join(c.Dir, file.Name()))
	}
	return bundle
}
