package defender

import (
	"testing"
	r "github.com/neurocollective/go_result/result"
	"fmt"
	"errors"
)

func GetServiceConfig(serviceName string) *r.Result {
	serviceConfigs := map[string]map[string]string {
		"testService": map[string]string{
			"url": "http://test.aexp.com/whatevs",
			"content": "application/json",
		},
	}

	serviceConfig, present := serviceConfigs[serviceName]
	if !present {
		return &r.Result{ nil, errors.New(serviceName + " not found!") }
	}
	return &r.Result{ serviceConfig, nil }
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

func TestDefender(t *testing.T) {

	result := New().Next(func(result *r.Result) *r.Result {

			return GetServiceConfig("testService")
		}).Next(func(result *r.Result) *r.Result {

			serviceConfig, ok := result.Value.(map[string]string)

			if !ok {
				return &Result { nil, errors.New("failed to assert `map[string]string`!")}
			}

			response := MockApiCall(serviceConfig)
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