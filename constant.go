package ginger

import "github.com/metadiv-io/ginger/internal/constant"

// expose constants

const (
	ERR_CODE_UNAUTHORIZED          = constant.ERR_CODE_UNAUTHORIZED
	ERR_CODE_FORBIDDEN             = constant.ERR_CODE_FORBIDDEN
	ERR_CODE_INTERNAL_SERVER_ERROR = constant.ERR_CODE_INTERNAL_SERVER_ERROR
)

const (
	LOCALE_EN  = constant.LOCALE_EN
	LOCALE_ZHT = constant.LOCALE_ZHT
	LOCALE_ZHS = constant.LOCALE_ZHS
)

const (
	HEADER_TRACE_ID = constant.HEADER_TRACE_ID
	HEADER_LOCALE   = constant.HEADER_LOCALE
	HEADER_TRACE    = constant.HEADER_TRACE
)
