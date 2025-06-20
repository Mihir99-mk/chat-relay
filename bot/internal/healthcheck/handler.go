package healthcheck

import (
	"bot/config"
	sql "bot/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheckHandler(config config.IConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := sql.GetSqlDriver(config.Env())
		if err != nil {
			return err
		}

		if err := db.DB().PingContext(c.Request().Context()); err != nil {
			return c.JSON(http.StatusServiceUnavailable, echo.Map{
				"status":  "error",
				"message": "Database unreachable",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"status":  "ok",
			"message": "Service and DB are healthy",
		})
	}
}
