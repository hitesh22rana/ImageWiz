package handlers

import (
	"io"
	"net/http"

	"github.com/hitesh22rana/imagewiz/tasks"
	"github.com/labstack/echo/v4"
)

func ResizeImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Upload failed")
	}

	fileData, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Failed to open the file")
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Failed to read the file")
	}

	resizeTasks, err := tasks.NewImageResizeTasks(data, file.Filename)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Could not create image resize tasks")
	}

	client := tasks.GetClient()
	for _, task := range resizeTasks {
		if _, err = client.Enqueue(task); err != nil {
			return echo.NewHTTPError(echo.ErrBadRequest.Code, "Could not enqueue image resize task")
		}
	}

	return c.JSON(http.StatusOK, "Image uploaded and resizing tasks started")
}
