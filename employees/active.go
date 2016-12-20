package employees

const (
	Active    ActiveFilter = "active"
	NotActive ActiveFilter = "notActive"
	All       ActiveFilter = "all"
)

type ActiveFilter string
