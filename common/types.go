package common

func IntPtr(v int) *int {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

func StringValues(ptrs []*string) []string {
	values := make([]string, len(ptrs))
	for i, p := range ptrs {
		if p != nil {
			values[i] = *p
		}
	}
	return values
}

func BoolPtr(v bool) *bool {
	return &v
}
