package qs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xackery/req-web/db"
)

func QsGet(c echo.Context) error {
	var err error
	args := map[string]interface{}{}

	type Qs struct {
		ID        int
		CharID    int `db:"char_id"`
		Event     int
		EventDesc string `db:"event_desc"`
		Time      int
	}
	var result = struct {
		CharID int
		Qs     []*Qs
	}{}

	args["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	query := `SELECT
	id, char_id, event, event_desc, time
	FROM qs_player_events
	WHERE char_id = :id
	ORDER BY time DESC
	LIMIT 1000`

	rows, err := db.Instance.NamedQueryContext(c.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("query qs: %w", err)
	}

	fmt.Println("foo")
	for rows.Next() {
		q := Qs{}
		err = rows.StructScan(&q)
		if err != nil {
			return fmt.Errorf("query qs scan: %w", err)
		}
		fmt.Println(q)
		result.Qs = append(result.Qs, &q)
	}
	return c.Render(http.StatusOK, "qs_get", result)
}
