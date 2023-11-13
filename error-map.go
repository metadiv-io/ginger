package ginger

const (
	locale_default = "default"
)

var errMapObj = errMap{}

type errMap map[string]map[string]string

func RegisterError(code, message string, locale ...string) {
	if len(locale) == 0 {
		locale = append(locale, locale_default)
	}
	if _, ok := errMapObj[locale[1]]; !ok {
		errMapObj[locale[1]] = make(map[string]string)
	}
	errMapObj[locale[1]][code] = message
}

func GetError(code string, locale ...string) string {
	if len(locale) == 0 {
		locale = append(locale, locale_default)
	}
	if _, ok := errMapObj[locale[1]]; !ok {
		return ""
	}
	if _, ok := errMapObj[locale[1]][code]; !ok {
		return ""
	}
	return errMapObj[locale[1]][code]
}
