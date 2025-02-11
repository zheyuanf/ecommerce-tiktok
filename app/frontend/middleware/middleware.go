package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterMiddleware(h *server.Hertz) {
	h.Use(GlobalAuth())
}

func RegisterJWTMiddleware(h *server.Hertz) {
	h.Use(GlobalJWTAuth())
}
