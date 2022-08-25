package utils

func StringPtr(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}
