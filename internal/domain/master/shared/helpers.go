package shared

// ValueOrEmpty returns "" if nil
func ValueOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Ptr returns pointer to value
func Ptr[T any](v T) *T {
	return &v
}
