package item

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
)

func ItemGet(c echo.Context) error {
	var err error
	args := map[string]interface{}{}
	var item = struct {
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
	FROM items
	WHERE id = :id`

	rows, err := db.Instance.NamedQueryContext(c.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("query item: %w", err)
	}

	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			return fmt.Errorf("query item scan: %w", err)
		}
		item.ShortNames = SlotShortNames(item.Slots)
	}
	return c.Render(http.StatusOK, "item_get", item)
}
