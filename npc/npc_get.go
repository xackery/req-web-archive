package npc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
)

func NpcGet(c echo.Context) error {
	var err error
	args := map[string]interface{}{}
	var npc = struct {
		ID         int
		Name       string
		AC         int
		Slots      int
		ShortNames string
		Weight     string
	}{}

	args["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	query := `SELECT
	id, name, ac, slots
	FROM npcs
	WHERE id = :id`

	rows, err := db.Instance.NamedQueryContext(c.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("query npc: %w", err)
	}

	for rows.Next() {
		err = rows.StructScan(&npc)
		if err != nil {
			return fmt.Errorf("query npc scan: %w", err)
		}
	}
	return c.Render(http.StatusOK, "npc_get", npc)
}
