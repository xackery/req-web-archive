package spell

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
)

func SpellGet(c echo.Context) error {
	var err error
	args := map[string]interface{}{}
	var spell = struct {
		ID        int
		Name      string
		Cast_time int
		Mana      int
		Range     int
	}{}

	args["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	query := "SELECT id, name, cast_time, mana, `range` FROM spells_new WHERE id = :id"

	rows, err := db.Instance.NamedQueryContext(c.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("query spell: %w", err)
	}

	for rows.Next() {
		err = rows.StructScan(&spell)
		if err != nil {
			return fmt.Errorf("query spell scan: %w", err)
		}
	}
	return c.Render(http.StatusOK, "spell_get", spell)
}
