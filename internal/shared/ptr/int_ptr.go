package ptr

// Int32OrNil returns nil if the value is zero.
func Int32OrNil(v *int32) *int32 {
	if v == nil || *v == 0 {
		return nil
	}
	return v
}
