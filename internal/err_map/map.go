package err_map

import "github.com/metadiv-io/ginger/internal/constant"

var ErrMap = errMap{}

type errMap map[string]map[string]string

func (m *errMap) Register(code string, locale string, message string) {
	if _, ok := (*m)[locale]; !ok {
		(*m)[locale] = make(map[string]string)
	}
	(*m)[locale][code] = message
}

func (m *errMap) Get(code string, locale string) string {
	if locale == "" {
		locale = constant.LOCALE_EN
	}
	if _, ok := (*m)[locale]; !ok {
		return ""
	}
	if _, ok := (*m)[locale][code]; !ok {
		return ""
	}
	return (*m)[locale][code]
}

func init() {
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_EN, "Unauthorized")
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_ZHT, "未經授權 (Unauthorized)")
	ErrMap.Register(constant.ERR_CODE_UNAUTHORIZED, constant.LOCALE_ZHS, "未经授权 (Unauthorized)")

	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_EN, "Forbidden")
	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_ZHT, "禁止 (Forbidden)")
	ErrMap.Register(constant.ERR_CODE_FORBIDDEN, constant.LOCALE_ZHS, "禁止 (Forbidden)")

	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_EN, "Internal Server Error")
	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_ZHT, "內部伺服器錯誤")
	ErrMap.Register(constant.ERR_CODE_INTERNAL_SERVER_ERROR, constant.LOCALE_ZHS, "内部服务器错误")
}
