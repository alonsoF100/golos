package service

const (
	maxLimit     = 100
	defaultLimit = 20
)

func validateLimit(limit int) int {
	if limit <= 0 {
		return defaultLimit
	}
	if limit > maxLimit {
		return maxLimit
	}
	return limit
}

func validateOffset(offset int) int {
	if offset < 0 {
		return 0
	}

	return offset
}
