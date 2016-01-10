package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSoftware_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/software", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":1,"name":"mysql"},
            {"id":2,"name":"postgresql"}]}`)
	})

	software, _, err := client.Software.List()
	if err != nil {
		t.Errorf("Software.List returned error: %v", err)
	}

	expected := []Software{
		{ID: 1, Name: "mysql"},
		{ID: 2, Name: "postgresql"},
	}

	if !reflect.DeepEqual(software, expected) {
		t.Errorf("Software.List returned\n %+v,\n expected\n %+v",
			software, expected)
	}
}
