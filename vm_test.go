package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVM_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, ``)
	})

	vms, _, err := client.VM.List()
	if err != nil {
		t.Errorf("VM.List returned error: %v", err)
	}

	expected := []VM{
		{
			AccessUsers:  []string{},
			BackupPolicy: BackupPolicy{},
			Bandwidth:    100,
		},
		//     		CPU             int         `json:"cpu"`
		//     		CurrentSnapshot interface{} `json:"currentSnapshot"`
		//     		Disk            int         `json:"disk"`
		//     		Distribution    struct {
		//     			Bitness     int    `json:"bitness"`
		//     			Description string `json:"description"`
		//     			ID          int    `json:"id"`
		//     			Name        string `json:"name"`
		//     		} `json:"distribution"`
		//     		ID               int           `json:"id"`
		//     		InternalID       string        `json:"internalId"`
		//     		IPAddresses      []interface{} `json:"ipAddresses"`
		//     		Memory           int           `json:"memory"`
		//     		Name             string        `json:"name"`
		//     		OwnerUser        interface{}   `json:"ownerUser"`
		//     		PrivateIPAddress string        `json:"privateIpAddress"`
		//     		PublicKeys       []struct {
		//     			ID        int         `json:"id"`
		//     			Installed bool        `json:"installed"`
		//     			Name      string      `json:"name"`
		//     			OwnerUser interface{} `json:"ownerUser"`
		//     			PublicKey interface{} `json:"publicKey"`
		//     			TimeAdded interface{} `json:"timeAdded"`
		//     			Type      string      `json:"type"`
		//     		} `json:"publicKeys"`
		//     		State     string `json:"state"`
		//     		TariffID  int    `json:"tariffId"`
		//     		TimeAdded int    `json:"timeAdded"`
		// },
	}

	if !reflect.DeepEqual(vms, expected) {
		t.Errorf("VM.List returned\n %+v,\n expected\n %+v",
			vms, expected)
	}
}
