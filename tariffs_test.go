package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTariffs_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/tariffs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":3,"name":"Облачный 512",
            "memory":512,"disk":16384,"cpu":1,"ipCount":1,"onDemand":false,
            "forWindows":false},{"id":1,"name":"Облачный без границ",
            "memory":512,"disk":8000,"cpu":0,"ipCount":0,"onDemand":true,
            "forWindows":false},{"id":7,"name":"Облачный 1024 для Windows",
            "memory":1024,"disk":32768,"cpu":1,"ipCount":1,"onDemand":false,
            "forWindows":true},{"id":2,"name":"Облачный без границ для Windows",
            "memory":512,"disk":8000,"cpu":0,"ipCount":0,"onDemand":true,
            "forWindows":true}]}`)
	})

	tariffs, _, err := client.Tariffs.List()
	if err != nil {
		t.Errorf("Tariffs.List returned error: %v", err)
	}

	expected := []Tariff{
		{
			ID:         3,
			Name:       "Облачный 512",
			Memory:     512,
			Disk:       16384,
			CPU:        1,
			IPCount:    1,
			OnDemand:   false,
			ForWindows: false,
		},
		{
			ID:         1,
			Name:       "Облачный без границ",
			Memory:     512,
			Disk:       8000,
			CPU:        0,
			IPCount:    0,
			OnDemand:   true,
			ForWindows: false,
		},
		{
			ID:         7,
			Name:       "Облачный 1024 для Windows",
			Memory:     1024,
			Disk:       32768,
			CPU:        1,
			IPCount:    1,
			OnDemand:   false,
			ForWindows: true,
		},
		{
			ID:         2,
			Name:       "Облачный без границ для Windows",
			Memory:     512,
			Disk:       8000,
			CPU:        0,
			IPCount:    0,
			OnDemand:   true,
			ForWindows: true,
		},
	}

	if !reflect.DeepEqual(tariffs, expected) {
		t.Errorf("Tariffs.List returned\n %+v,\n expected\n %+v", tariffs, expected)
	}
}
