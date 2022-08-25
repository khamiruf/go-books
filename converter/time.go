package converter

import (
	"database/sql"
	"time"
)

func SqlToTime(time *sql.NullTime) *time.Time {
	if !time.Valid {
		return nil
	}

	return &time.Time
}
