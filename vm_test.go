package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestVM_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"OK","result":[{"id":88534,"name":"win server",
			"internalId":"vm88534","memory":1024,"disk":24576,"cpu":1,
			"bandwidth":100,"tariffId":20,"ipAddresses":[],
			"privateIpAddress":"10.7.13.159","state":"VIR_DOMAIN_SHUTOFF",
			"timeAdded":1412089961298,"currentSnapshot":null,"backupPolicy":
			{"quantity":1,"frequency":72},"distribution":{"id":219,"name":
			"WINDOWS_2008_R2","description":"Win Server 2008 R2 x64",
			"bitness":64},"ownerUser":null,"accessUsers":[],"publicKeys":
			[{"id":2,"name":"1","type":"DSA","publicKey":null,"timeAdded":null,
			"installed":false,"ownerUser":null}]}]}`)
	})

	vms, _, err := client.VM.List()
	if err != nil {
		t.Errorf("VM.List returned error: %v", err)
	}

	expected := []VM{
		{
			ID:               88534,
			Name:             "win server",
			InternalID:       "vm88534",
			Memory:           1024,
			Disk:             24576,
			CPU:              1,
			Bandwidth:        100,
			TariffID:         20,
			IPAddresses:      []string{},
			PrivateIPAddress: "",
			State:            "VIR_DOMAIN_SHUTOFF",
			TimeAdded:        Timestamp{time.Unix(1412089961298, 0)},
			CurrentSnapshot:  nil,
			BackupPolicy:     BackupPolicy{1, 72},
			Distribution:     Distribution{219, "WINDOWS_2008_R2", "Win Server 2008 R2 x64", 64},
			OwnerUser:        nil,
			AccessUsers:      []string{},
			PublicKeys:       []VMPublicKey{{2, "1", "DSA", nil, nil, false, nil}},
		},
	}

	if !reflect.DeepEqual(vms, expected) {
		t.Errorf("VM.List returned\n %+v,\n expected\n %+v",
			vms, expected)
	}
}
