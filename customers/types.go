package customers

type QueryMatchType string

const (
	Contains 	QueryMatchType = "contains"
	StartsWith 	QueryMatchType = "starts_with"
)
