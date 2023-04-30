package assert

import (
	"github.com/neurocollective/go_result/result"
	"errors"
)

type Result = result.Result

type TypedResult[V any] struct {
	Value V
	Error error
}

func (r *TypedResult[V]) Ok() bool {
	if r == nil || r.Error != nil {
		return false
	}
	return true
}

func GetTypedAction[T any](
	action func(*TypedResult[T]) *Result,
) func(*Result) *Result {

	return func(result *Result) *Result {

		assertedValue, ok := result.Value.(T)

		if !ok {
			return &Result{ nil, errors.New("could not assert a result.Value as requested") }
		}
		return action(&TypedResult[T]{ assertedValue, nil })
	}
}

