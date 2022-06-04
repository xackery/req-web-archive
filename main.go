package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xackery/req-web/db"
	"github.com/xackery/req-web/item"
	"github.com/xackery/req-web/npc"
	"github.com/xackery/req-web/spell"
	etemplate "github.com/xackery/req-web/template"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed:", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}

	t := &etemplate.Template{
		Templates: template.Must(template.ParseGlob("web/template/*.html")),
	}

	e := echo.New()
	e.HideBanner = true
	e.Debug = true
	e.Server.ReadTimeout = 5 * time.Second
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", root)
	e.GET("/item/:id", item.ItemGet)
	e.GET("/spell/:id", spell.SpellGet)
	e.GET("/npc/:id", npc.NpcGet)
	e.Static("/static", "web/static")
	fmt.Println("http://localhost:8080")
	e.Renderer = t

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}

func root(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
