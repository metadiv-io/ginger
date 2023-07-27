package caller

import "github.com/gin-gonic/gin"

func GET[T any](ctx *gin.Context, host, path string,
	params map[string]string, headers map[string]string) (*Response[T], error) {
	return get[T](ctx, host, path, params, headers)
}

func POST[T any](ctx *gin.Context, host, path string,
	body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, host, path, "POST", body, headers)
}

func PUT[T any](ctx *gin.Context, host, path string,
	body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, host, path, "PUT", body, headers)
}

func DELETE[T any](ctx *gin.Context, host, path string,
	body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, host, path, "DELETE", body, headers)
}
