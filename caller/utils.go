package caller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func get[T any](ctx *gin.Context, host string, path string,
	params map[string]string, headers map[string]string) (*Response[T], error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Authorization"] = getAuthToken(ctx)

	path += "?"
	for k, v := range params {
		path += k + "=" + v + "&"
	}
	path = path[:len(path)-1]

	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("failed to unmarshal response body: ", string(bodyBytes))
		return nil, err
	}
	return &response, nil
}

func nonGet[T any](ctx *gin.Context, host, path, method string,
	body interface{}, headers map[string]string) (*Response[T], error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Authorization"] = getAuthToken(ctx)

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, host+path, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println("failed to unmarshal response body: ", string(bodyBytes))
		return nil, err
	}
	return &response, nil
}

func getAuthToken(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	token := ctx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}
