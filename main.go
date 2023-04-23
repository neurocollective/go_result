package go_result

type Result[V any] struct {
	Value V
	Error error
}

func (r *Result[V]) Ok() bool {
	if r.Error != nil {
		return false
	}
	return true
}

func (r *Result[V]) IsNil() bool {
	if r.Value == nil {
		return true
	}
	return false
}

func (r *Result[V]) Error() error {
	return r.Error
}

func (r *Result[V]) Value() V {
	return r.Value
}

type Janitor[V any] struct {
	CurrentResult *Result[V]
	Handlers []func(Result[V])
}

func (j *Janitor[V]) SetHandlers(handlers ...func(Result[V])) *Janitor[V] {
	j.Handlers = handlers
	return j
}

func (r *Janitor[V]) Receive(newResult *Result[V]) V {
	r.CurrentResult = newResult

	if len(r.Handlers) > 0 {
		for _, handler := range handlers {
			if handler == nil {
				continue
			}
			handler(newResult)
		}
	}
	return r.CurrentResult.Value
}

func NewJanitor[V any](handlers ...func(Result[V])) *Janitor[V] {
	janitor := Janitor[V]{}
	if len(handlers) > 0 {
		janitor.SetHandlers(handlers)
	}
	return &janitor
}
