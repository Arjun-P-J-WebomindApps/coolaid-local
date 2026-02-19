package sqlnull

import "database/sql"

func Float64(v *float64) sql.NullFloat64 {
	if v == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *v,
		Valid:   true,
	}
}

func Float64Ptr(v sql.NullFloat64) *float64 {
	if !v.Valid {
		return nil
	}
	return &v.Float64
}

func Float64Value(v sql.NullFloat64) float64 {
	if !v.Valid {
		return 0.0
	}
	return v.Float64
}
