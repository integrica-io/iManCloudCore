package library

type PagingMode string
type SortOrder string

const (
	Standard 		PagingMode = "standard"
	StandardCursor	PagingMode = "standard_cursor"	
)

const (
	Asc		SortOrder = "asc"
	Desc	SortOrder = "desc"
)
