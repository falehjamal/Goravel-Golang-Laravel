package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		// return ctx.Response().View().Make("welcome.tmpl", map[string]any{
		// 	"version": support.Version,
		// })
		return ctx.Response().Json(http.StatusOK, http.Json{
			"pesan": "ini Goravel alias Laravel tapi pakai bahasa Go",
		  })
	})
}
