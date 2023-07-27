package ginger

var ErrMap = &errMap{}

type errMap map[string]errMsg

type errMsg map[string]string

func (m *errMap) Register(code string, locale string, message string) {
	if _, ok := (*m)[locale]; !ok {
		(*m)[locale] = make(errMsg)
	}
	(*m)[locale][code] = message
}

func (m *errMap) Get(code string, locale string) string {
	if locale == "" {
		locale = LOCALE_EN
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
	ErrMap.Register(ERR_CODE_UNAUTHORIZED, LOCALE_EN, "Unauthorized")
	ErrMap.Register(ERR_CODE_UNAUTHORIZED, LOCALE_ZHT, "未經授權 (Unauthorized)")
	ErrMap.Register(ERR_CODE_UNAUTHORIZED, LOCALE_ZHS, "未经授权 (Unauthorized)")

	ErrMap.Register(ERR_CODE_FORBIDDEN, LOCALE_EN, "Forbidden")
	ErrMap.Register(ERR_CODE_FORBIDDEN, LOCALE_ZHT, "禁止 (Forbidden)")
	ErrMap.Register(ERR_CODE_FORBIDDEN, LOCALE_ZHS, "禁止 (Forbidden)")

	ErrMap.Register(ERR_CODE_INTERNAL_SERVER_ERROR, LOCALE_EN, "Internal Server Error")
	ErrMap.Register(ERR_CODE_INTERNAL_SERVER_ERROR, LOCALE_ZHT, "內部伺服器錯誤")
	ErrMap.Register(ERR_CODE_INTERNAL_SERVER_ERROR, LOCALE_ZHS, "内部服务器错误")
}
