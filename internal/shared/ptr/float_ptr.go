package ptr

// Float64Value returns the dereferenced value.
// If the pointer is nil, it returns 0.
func Float64Value(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}
