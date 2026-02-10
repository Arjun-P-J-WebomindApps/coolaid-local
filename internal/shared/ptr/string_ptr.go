package ptr

func String(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func StringPtr(p *string) *string {
	if p == nil {
		return nil
	}
	v := *p
	return &v
}

func StringOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
