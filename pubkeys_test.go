package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestPubKeys_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/pubkeys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":254,"name":"my-key2",
            "type":"DSA","publicKey":"ssh-dss AAAAB3NzaC1kc3MAAACBAK5uLwicCrFEpaV...Nt0Q7P45rZjNNTag2c= user@host",
            "timeAdded":1414156464315},{"id":2,"name":"my-key1","type":"DSA",
            "publicKey":"ssh-dss AAAAB3NzaC1kc3MAAACBAK5uLwicCrFEpaV...Nt0Q7P45rZjNNTag2c= user@host",
            "timeAdded":1351093292416}]}`)
	})

	pubkeys, _, err := client.PubKeys.List()
	if err != nil {
		t.Errorf("PubKeys.List returned error: %v", err)
	}

	expected := []PubKey{
		{
			ID:        254,
			Name:      "my-key2",
			Type:      "DSA",
			PublicKey: "ssh-dss AAAAB3NzaC1kc3MAAACBAK5uLwicCrFEpaV...Nt0Q7P45rZjNNTag2c= user@host",
			TimeAdded: Timestamp{time.Unix(1414156464315, 0)},
		},
		{
			ID:        2,
			Name:      "my-key1",
			Type:      "DSA",
			PublicKey: "ssh-dss AAAAB3NzaC1kc3MAAACBAK5uLwicCrFEpaV...Nt0Q7P45rZjNNTag2c= user@host",
			TimeAdded: Timestamp{time.Unix(1351093292416, 0)},
		},
	}

	if !reflect.DeepEqual(pubkeys, expected) {
		t.Errorf("PubKeys.List returned\n %+v,\n expected\n %+v",
			pubkeys, expected)
	}
}
