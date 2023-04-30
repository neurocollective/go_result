package assert

import (
	"testing"
	"github.com/neurocollective/go_result/defender"
	"fmt"
	"errors"
)

func GetServiceConfig(serviceName string) *Result {
	serviceConfigs := map[string]map[string]string {
		"testService": map[string]string{
			"url": "http://test.aexp.com/whatevs",
			"content": "application/json",
		},
	}

	serviceConfig, present := serviceConfigs[serviceName]
	if !present {
		return &Result{ nil, errors.New(serviceName + " not found!") }
	}
	return &Result{ serviceConfig, nil }
}

func MockApiCall(config map[string]string) map[string]string {
	return map[string]string {
		"card": "platinum",
		"offer": "false",
	}
}

func DoThingTwo(response map[string]string) *Result {
	return &Result{ nil, nil }
}

func HandleError(result *Result) *Result {
	fmt.Println(result.Error)
	return nil
}

func TestDefenderWithAssert(t *testing.T) {

	result := defender.New().Next(func(result *Result) *Result {

			return GetServiceConfig("testService")
		}).Next(GetTypedAction[map[string]string](
			func(result *TypedResult[map[string]string]) *Result {

				response := MockApiCall(result.Value)
				return DoThingTwo(response)
			}),
		).Error(func(result *Result) *Result {

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