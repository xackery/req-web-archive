package zone

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/xackery/req-web/db"
)

var (
	mu             sync.RWMutex
	zoneShortNames = make(map[int]string)
)

// ZoneShortNameByID returns a zone's short name
func ZoneShortNameByID(id int) (string, error) {
	var err error
	mu.Lock()
	defer mu.Unlock()

	name, ok := zoneShortNames[id]
	if ok {
		return name, nil
	}
	name, err = zoneShortName(id)
	if err != nil {
		return "", fmt.Errorf("zoneName: %w", err)
	}
	zoneShortNames[id] = name
	return name, nil
}

func zoneShortName(id int) (string, error) {
	args := map[string]interface{}{}
	args["id"] = id

	query := `SELECT
	short_name
	FROM zone
	WHERE id = :id`

	var zone = struct {
		Short_name sql.NullString
	}{}

	rows, err := db.Instance.NamedQueryContext(context.Background(), query, args)
	if err != nil {
		return "", fmt.Errorf("query zone: %w", err)
	}

	for rows.Next() {
		err = rows.StructScan(&zone)
		if err != nil {
			return "", fmt.Errorf("query zone scan: %w", err)
		}
		if zone.Short_name.Valid {
			return zone.Short_name.String, nil
		}
	}
	return "", fmt.Errorf("unknown zone: %d", id)
}
