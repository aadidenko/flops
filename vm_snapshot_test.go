package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestVMOperation_Snapshots(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/snapshots", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":2206,"name":"my-snapshot",
            "description":"My first snapshot","parentSnapshotId":null,
            "timeAdded":1414156778188}]}`)
	})

	snapshots, _, err := client.VM.Snapshots(1)
	if err != nil {
		t.Errorf("VM.Snapshots returned error: %v", err)
	}

	expected := []Snapshot{
		{
			ID:               2206,
			Name:             "my-snapshot",
			Description:      "My first snapshot",
			ParentSnapshotID: nil,
			TimeAdded:        Timestamp{time.Unix(1414156778188, 0)},
		},
	}

	if !reflect.DeepEqual(snapshots, expected) {
		t.Errorf("VM.Snapshots returned\n %+v,\n expected\n %+v",
			snapshots, expected)
	}
}
