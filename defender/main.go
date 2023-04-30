package defender

import (
	r "github.com/neurocollective/go_result/result"
	"fmt"
)

type Defender struct {
	CurrentResult *r.Result
	handlers []func(*r.Result) *r.Result
	errorHandler func(*r.Result) *r.Result
}

func New() *Defender {
	defender := Defender{}
	return &defender
}

func (d *Defender) Next(action func(*r.Result) *r.Result) *Defender {
	d.handlers = append(d.handlers, action)
	return d
}

func (d *Defender) Error(errorAction func(*r.Result) *r.Result) *Defender {
	d.errorHandler = errorAction
	return d
}

func (d *Defender) Run() *r.Result {
	var lastResult *r.Result
	for _, action := range d.handlers {
		currentResult := action(lastResult)

		if !currentResult.Ok() {  
			return d.errorHandler(currentResult)
		}
		lastResult = currentResult
	}
	return &r.Result{ "done", nil }
}

