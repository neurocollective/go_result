package result

import (
	"testing"
)

func DoAThing() *Result {
	return &Result{ "super valuabale data!", nil }
}

func DoADifferentThing() *Result {
	return &Result{ nil, nil }
}

func TestNewResult(t *testing.T) {

	result := DoAThing()

	if result.Ok() {
		t.Log("result is:", result.Value)
	} else {
		t.Errorf("oh no: %s", result.Error)
		return
	}

	resultTwo := DoADifferentThing()

	if result.Ok() {
		t.Log("result 2 is:", resultTwo.Value)
		t.Log("result 2 is nil:", resultTwo.Value == nil)
	} else {
		t.Errorf("oh no: %s", resultTwo.Error)
	}
}

