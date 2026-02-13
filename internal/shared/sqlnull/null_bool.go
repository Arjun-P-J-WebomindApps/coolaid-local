package sqlnull

import "database/sql"

func Bool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true,
	}
}

func BoolValue(nb sql.NullBool) bool {
	if !nb.Valid {
		return false
	}
	return nb.Bool
}
