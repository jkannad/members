package models

type error struct {
	err map[string]string
}

func (e *error) Add(field, value string) {
	e.err[field] = value
}

func (e *error) Get(field string) string {
	return e.err[field]
}
