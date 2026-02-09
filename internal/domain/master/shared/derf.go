package shared

func Deref[T any](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}
