package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOperation_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/operation/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":{"id":1,"vmId":85662,
            "status":"DONE","operationType":"VM_DESTROY","percentage":100,
            "errorMessage":null,"errorCode":null}}`)
	})

	operation, _, err := client.Operation.Get(1)
	if err != nil {
		t.Errorf("Operation.Get returned error: %v", err)
	}

	expected := &Operation{
		ID:           1,
		VMID:         85662,
		Status:       "DONE",
		Type:         "VM_DESTROY",
		Percentage:   100,
		ErrorMessage: nil,
		ErrorCode:    nil,
	}

	if !reflect.DeepEqual(operation, expected) {
		t.Errorf("Operation.Get returned\n %+v,\n expected\n %+v",
			operation, expected)
	}
}
