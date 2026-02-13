package ptr

// Int32OrNil returns nil if the value is zero.
func Int32OrNil(v *int32) *int32 {
	if v == nil || *v == 0 {
		return nil
	}
	return v
}

// Int32Value returns the dereferenced value.
// If the pointer is nil, it returns 0.
func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}
