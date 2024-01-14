package routes

import (
	"github.com/hitesh22rana/imagewiz/handlers"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.POST("/api/v1/resize-image", handlers.ResizeImage)
}
