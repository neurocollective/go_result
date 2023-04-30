package result

type Result struct {
	Value any
	Error error
}

func As[V any](result *Result) V {
	castValue := result.Value.(V)
	return castValue
}

func (r *Result) Ok() bool {
	if r == nil || r.Error != nil {
		return false
	}
	return true
}
