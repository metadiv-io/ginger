package ginger

import (
	"github.com/metadiv-io/ginger/internal/constant"
	"github.com/metadiv-io/ginger/internal/err_map"
)

func RegisterError(code, locale, message string) {
	if _, ok := err_map.ErrMap[locale]; !ok {
		err_map.ErrMap[locale] = make(map[string]string)
	}
	err_map.ErrMap[locale][code] = message
}

func GetError(code, locale string) string {
	if locale == "" {
		locale = constant.LOCALE_EN
	}
	if _, ok := err_map.ErrMap[locale]; !ok {
		return ""
	}
	if _, ok := err_map.ErrMap[locale][code]; !ok {
		return ""
	}
	return err_map.ErrMap[locale][code]
}
