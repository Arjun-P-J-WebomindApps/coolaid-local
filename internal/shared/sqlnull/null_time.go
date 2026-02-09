package sqlnull

import (
	"database/sql"
	"time"
)

// TimePtr converts sql.NullTime to *time.Time.
// Returns nil if the value is NULL.
func TimePtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
