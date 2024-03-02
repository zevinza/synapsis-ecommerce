package lib

func RevInt64(i *int64) int64 {
	if nil == i {
		return 0
	}
	return *i
}

func RevStr(s *string) string {
	if nil == s {
		return ""
	}
	return *s
}

func RevFloat64(f *float64) float64 {
	if nil == f {
		return 0
	}
	return *f
}
