package result

type Result struct {
	Value any
	Error error
}

func (r *Result) Ok() bool {
	if r == nil || r.Error != nil {
		return false
	}
	return true
}
