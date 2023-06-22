package errorConstants

type CommonStruct struct {
	ErrorName string
	Message   string
}

type DuplicateStruct struct {
	Field       string
	Value       string
	Description string
}
