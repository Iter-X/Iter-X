package interceptor

import (
	"context"
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	xerr "github.com/protoc-gen/protoc-gen-go-errors/errors"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/iter-x/iter-x/internal/common/cnst"
)

func I18n(bundle *i18n.Bundle) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = withLanguage(ctx, md)
		}

		resp, err := handler(ctx, req)
		if err != nil {
			var xErr *xerr.Error
			if !errors.As(err, &xErr) {
				return nil, err
			}
			if !ok {
				return nil, err
			}
			localize, err2 := i18n.NewLocalizer(bundle,
				getHeader(md, cnst.HttpHeaderLang), getHeader(md, cnst.HttpHeaderAcceptLang)).
				Localize(&i18n.LocalizeConfig{MessageID: xErr.GetReason()})
			if err2 == nil {
				xErr.Message = localize
			}
			return nil, err
		}
		return resp, nil
	}
}

var matcher = language.NewMatcher([]language.Tag{
	language.MustParse(cnst.LangEn),
	language.MustParse(cnst.LangZhCn),
})

// withLanguage determines the language from metadata and sets it in context
func withLanguage(ctx context.Context, md metadata.MD) context.Context {
	lang := getHeader(md, cnst.HttpHeaderLang)
	if lang == "" {
		lang = getHeader(md, cnst.HttpHeaderAcceptLang)
	}

	if lang == "" {
		return context.WithValue(ctx, cnst.CtxKeyLang, cnst.LangEn)
	}

	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil || len(tags) == 0 {
		return context.WithValue(ctx, cnst.CtxKeyLang, cnst.LangEn)
	}

	tag, _, _ := matcher.Match(tags[0])
	if tag == language.MustParse(cnst.LangZhCn) {
		return context.WithValue(ctx, cnst.CtxKeyLang, cnst.LangZhCn)
	}
	return context.WithValue(ctx, cnst.CtxKeyLang, cnst.LangEn)
}

func getHeader(md metadata.MD, key string) string {
	if md == nil {
		return ""
	}
	vals := md.Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}
