package rename

import (
	"net/http"

	"github.com/labstack/echo"
)

func APITablesHandler(c echo.Context) error {

	tables := []string{
		"table1",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": tables, "error": false})
}
