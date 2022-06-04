package testdata

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
	etemplate "github.com/xackery/req-web/template"
)

func InitEcho() (*echo.Echo, echo.Context, *ResponseWriterWrapper, error) {
	err := db.Init(context.Background())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("db.Init: %w", err)
	}

	tmpl := &etemplate.Template{
		Templates: template.Must(template.ParseGlob("../web/template/*.html")),
	}
	e := echo.New()

	e.HideBanner = true
	e.Debug = true
	e.Renderer = tmpl
	w := &ResponseWriterWrapper{}
	c := e.NewContext(&http.Request{}, w)

	return e, c, w, nil
}

type ResponseWriterWrapper struct {
	//w          http.ResponseWriter
	Body       bytes.Buffer
	statusCode int
}

func (i *ResponseWriterWrapper) Write(buf []byte) (int, error) {
	i.Body.Write(buf)
	return 0, nil
	//return i.w.Write(buf)
}

func (i *ResponseWriterWrapper) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	//i.w.WriteHeader(statusCode)
}

func (i *ResponseWriterWrapper) Header() http.Header {
	return make(map[string][]string)
}
