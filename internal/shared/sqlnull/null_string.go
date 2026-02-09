package sqlnull

import "database/sql"

func String(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func StringPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

// sql.NullString → string (for reads)
func StringValueOrEmpty(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

func StringPtrOrEmpty(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	v := ns.String
	return &v
}
