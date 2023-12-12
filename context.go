package ginger

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/nanoid"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	Engine    IEngine
	GinCtx    *gin.Context
	Page      *sql.Pagination
	Sort      *sql.Sort
	Request   *T
	Response  *Response
	StartTime time.Time
	IsFile    bool
	IsResp    bool
}

func (ctx *Context[T]) GetEngine() IEngine {
	return ctx.Engine
}

func (ctx *Context[T]) SetEngine(engine IEngine) {
	ctx.Engine = engine
}

func (ctx *Context[T]) GetGinCtx() *gin.Context {
	return ctx.GinCtx
}

func (ctx *Context[T]) SetGinCtx(ginCtx *gin.Context) {
	ctx.GinCtx = ginCtx
}

func (ctx *Context[T]) GetPage() *sql.Pagination {
	return ctx.Page
}

func (ctx *Context[T]) SetPage(page *sql.Pagination) {
	ctx.Page = page
}

func (ctx *Context[T]) GetSort() *sql.Sort {
	return ctx.Sort
}

func (ctx *Context[T]) SetSort(sort *sql.Sort) {
	ctx.Sort = sort
}

func (ctx *Context[T]) GetRequest() *T {
	return ctx.Request
}

func (ctx *Context[T]) SetRequest(request *T) {
	ctx.Request = request
}

func (ctx *Context[T]) GetResponse() *Response {
	return ctx.Response
}

func (ctx *Context[T]) SetResponse(response *Response) {
	ctx.Response = response
}

func (ctx *Context[T]) GetStartTime() time.Time {
	return ctx.StartTime
}

func (ctx *Context[T]) SetStartTime(startTime time.Time) {
	ctx.StartTime = startTime
}

func (ctx *Context[T]) GetIsFile() bool {
	return ctx.IsFile
}

func (ctx *Context[T]) SetIsFile(isFile bool) {
	ctx.IsFile = isFile
}

func (ctx *Context[T]) GetIsResponded() bool {
	return ctx.IsResp
}

func (ctx *Context[T]) SetIsResponded(isResp bool) {
	ctx.IsResp = isResp
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx.Request.UserAgent()
}

func (ctx *Context[T]) BearerToken() string {
	token := ctx.GinCtx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func (ctx *Context[T]) OK(data any, page ...*sql.Pagination) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}
	// pagination
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}
	ctx.SetResponse(&Response{
		Success:    true,
		Pagination: pageResponse,
		Data:       data,
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
	})
	ctx.SetIsResponded(true)
}

func (ctx *Context[T]) OKFile(bytes []byte, filename ...string) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}

	var name string
	if len(filename) == 0 || filename[0] == "" {
		name = nanoid.NewSafe()
	} else {
		name = filename[0]
	}

	ctx.GinCtx.Header("Content-Disposition", "filename="+name)
	ctx.GinCtx.Data(http.StatusOK, ctx.determineFileType(name), bytes)
	ctx.IsResp = true
	ctx.IsFile = true
}

func (ctx *Context[T]) OKDownload(bytes []byte, filename ...string) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}

	var name string
	if len(filename) == 0 || filename[0] == "" {
		name = nanoid.NewSafe()
	} else {
		name = filename[0]
	}

	ctx.GinCtx.Header("Content-Disposition", "filename="+name)
	ctx.GinCtx.Data(http.StatusOK, "application/octet-stream", bytes)
	ctx.IsResp = true
	ctx.IsFile = true
}

func (ctx *Context[T]) Err(code string, locale ...string) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}
	ctx.SetResponse(&Response{
		Success:  false,
		Duration: time.Since(ctx.StartTime).Milliseconds(),
		Error: &Error{
			Code:    code,
			Message: GetError(code, locale...),
		},
	})
	ctx.SetIsResponded(true)
}

func (ctx *Context[T]) determineFileType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".aac":
		return "audio/aac"
	case ".abw":
		return "application/x-abiword"
	case ".arc":
		return "application/x-freearc"
	case ".avif":
		return "image/avif"
	case ".avi":
		return "video/x-msvideo"
	case ".azw":
		return "application/vnd.amazon.ebook"
	case ".bin":
		return "application/octet-stream"
	case ".bz":
		return "application/x-bzip"
	case ".bz2":
		return "application/x-bzip2"
	case ".cda":
		return "application/x-cdf"
	case ".csh":
		return "application/x-csh"
	case ".css":
		return "text/css"
	case ".csv":
		return "text/csv"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".eot":
		return "application/vnd.ms-fontobject"
	case ".epub":
		return "application/epub+zip"
	case ".gz":
		return "application/gzip"
	case ".gif":
		return "image/gif"
	case ".htm":
		return "text/html"
	case ".html":
		return "text/html"
	case ".ico":
		return "image/vnd.microsoft.icon"
	case ".ics":
		return "text/calendar"
	case ".jar":
		return "application/java-archive"
	case ".jpeg":
		return "image/jpeg"
	case ".jpg":
		return "image/jpeg"
	case ".js":
		return "text/javascript"
	case ".json":
		return "application/json"
	case ".jsonld":
		return "application/ld+json"
	case ".mid":
		return "audio/midi"
	case ".midi":
		return "audio/midi"
	case ".mjs":
		return "text/javascript"
	case ".mp3":
		return "audio/mpeg"
	case ".mp4":
		return "video/mp4"
	case ".mpeg":
		return "video/mpeg"
	case ".mpkg":
		return "application/vnd.apple.installer+xml"
	case ".odp":
		return "application/vnd.oasis.opendocument.presentation"
	case ".ods":
		return "application/vnd.oasis.opendocument.spreadsheet"
	case ".odt":
		return "application/vnd.oasis.opendocument.text"
	case ".oga":
		return "audio/ogg"
	case ".ogv":
		return "video/ogg"
	case ".ogx":
		return "application/ogg"
	case ".opus":
		return "audio/opus"
	case ".otf":
		return "font/otf"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	case ".php":
		return "application/x-httpd-php"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".rar":
		return "application/vnd.rar"
	case ".rtf":
		return "application/rtf"
	case ".sh":
		return "application/x-sh"
	case ".svg":
		return "image/svg+xml"
	case ".tar":
		return "application/x-tar"
	case ".tif":
		return "image/tiff"
	case ".tiff":
		return "image/tiff"
	case ".ts":
		return "video/mp2t"
	case ".ttf":
		return "font/ttf"
	case ".txt":
		return "text/plain"
	case ".vsd":
		return "application/vnd.visio"
	case ".wav":
		return "audio/wav"
	case ".weba":
		return "audio/webm"
	case ".webm":
		return "video/webm"
	case ".webp":
		return "image/webp"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".xhtml":
		return "application/xhtml+xml"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".xml":
		return "application/xml"
	case ".xul":
		return "application/vnd.mozilla.xul+xml"
	case ".zip":
		return "application/zip"
	case ".3gp":
		return "video/3gpp"
	case ".3g2":
		return "video/3gpp2"
	case ".7z":
		return "application/x-7z-compressed"
	default:
		return "application/octet-stream"
	}
}
