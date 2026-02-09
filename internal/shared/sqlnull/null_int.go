package sqlnull

import "database/sql"

func Int32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *i,
		Valid: true,
	}
}
