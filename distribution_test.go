package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDistribution_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/distribution", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":233,"name":"DEBIAN",
            "description":"Debian 7.6 x64","bitness":64},{"id":234,
            "name":"DEBIAN","description":"Debian 7.6 x86","bitness":32}]}`)
	})

	distribution, _, err := client.Distribution.List()
	if err != nil {
		t.Errorf("Distribution.List returned error: %v", err)
	}

	expected := []Distribution{
		{
			ID:          233,
			Name:        "DEBIAN",
			Description: "Debian 7.6 x64",
			Bitness:     64,
		},
		{
			ID:          234,
			Name:        "DEBIAN",
			Description: "Debian 7.6 x86",
			Bitness:     32,
		},
	}

	if !reflect.DeepEqual(distribution, expected) {
		t.Errorf("Distribution.List returned\n %+v,\n expected\n %+v",
			distribution, expected)
	}
}
