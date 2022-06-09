package zone

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
	"github.com/xackery/req-web/expansion"
)

func ZoneGet(c echo.Context) error {
	var err error
	args := map[string]interface{}{}
	var zone = struct {
		ID            int
		Long_name     string
		Short_name    sql.NullString
		ShortName     string
		Expansion     int
		ExpansionName string
		Min_level     int
		Safe_x        float32
		Safe_y        float32
	}{}

	args["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	query := `SELECT
	id, long_name, short_name, min_level, safe_x, safe_y, expansion
	FROM zone
	WHERE id = :id`

	rows, err := db.Instance.NamedQueryContext(c.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("query zone: %w", err)
	}

	for rows.Next() {
		err = rows.StructScan(&zone)
		if err != nil {
			return fmt.Errorf("query zone scan: %w", err)
		}
		if zone.Short_name.Valid {
			zone.ShortName = zone.Short_name.String
		} else {
			zone.ShortName = "unknown"
		}

		zone.ExpansionName, err = expansion.ExpansionNameByID(zone.Expansion - 1)
		if err != nil {
			return fmt.Errorf("ExpansionNameByID %d: %w", zone.Expansion-1, err)
		}
	}
	return c.Render(http.StatusOK, "zone_get", zone)
}
