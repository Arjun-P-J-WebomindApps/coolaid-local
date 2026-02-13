package sqlnull

import "database/sql"

func Int32(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *v,
		Valid: true,
	}
}

func Int32Ptr(v sql.NullInt32) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}
