package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestVMOperation_Backups(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/backups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":1,"size":123456,
            "timeAdded":1414156778188}]}`)
	})

	backups, _, err := client.VM.Backups(1)
	if err != nil {
		t.Errorf("VM.Backups returned error: %v", err)
	}

	expected := []Backup{
		{
			ID:        1,
			Size:      123456,
			TimeAdded: Timestamp{time.Unix(1414156778188, 0)},
		},
	}

	if !reflect.DeepEqual(backups, expected) {
		t.Errorf("VM.Backups returned\n %+v,\n expected\n %+v",
			backups, expected)
	}
}
