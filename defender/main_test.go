package defender

import (
	"testing"
	r "github.com/neurocollective/go_result/result"
	"fmt"
)

func DoThingOne(serviceName string) *r.Result {
	return &r.Result{ serviceName, nil }
}

func MockApiCall(config map[string]string) map[string]string {
	return map[string]string {
		"card": "platinum",
		"offer": "false",
	}
}

func DoThingTwo(response map[string]string) *r.Result {
	return &r.Result{ nil, nil }
}

func HandleError(result *r.Result) *r.Result {
	fmt.Println(result.Error)
	return nil
}

func TestNewResult(t *testing.T) {

	service := "nameOfAService"
	requestConfig := map[string]string {
		"url": "http://test.aexp.com/whatevs",
		"content": "application/json",
	}

	result := New().Next(func(result *r.Result) *r.Result {
			return DoThingOne(service)
		}).Next(func(result *r.Result) *r.Result {
			response := MockApiCall(requestConfig)
			return DoThingTwo(response)
		}).Error(func(result *r.Result) *r.Result {
			t.Errorf("oh noes, failed: %s", result.Error)
			return result
		}).Run()

	if result == nil {
		t.Errorf("oh no, nil result")
		return
	}

	if result.Ok() {
		t.Log("result is:", result.Value)
	} else {
		t.Errorf("oh no: %s", result.Error)
		return
	}
}