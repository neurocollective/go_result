package defender

import (
	"github.com/neurocollective/go_result/result"
)

type Result = result.Result

type Defender struct {
	CurrentResult *Result
	handlers []func(*Result) *Result
	errorHandler func(*Result) *Result
}

func New() *Defender {
	defender := Defender{}
	return &defender
}

func (d *Defender) Next(action func(*Result) *Result) *Defender {
	d.handlers = append(d.handlers, action)
	return d
}

func (d *Defender) Error(errorAction func(*Result) *Result) *Defender {
	d.errorHandler = errorAction
	return d
}

func (d *Defender) Run() *Result {
	var lastResult *Result
	for _, action := range d.handlers {
		currentResult := action(lastResult)

		if !currentResult.Ok() {  
			return d.errorHandler(currentResult)
		}
		lastResult = currentResult
	}
	return &Result{ "done", nil }
}

