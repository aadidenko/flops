package flops

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

func TestVMOperation_Rename(t *testing.T) {
	setup()
	defer teardown()

	newVMName := "New-VM-Name"

	mux.HandleFunc("/v1/vm/1/rename", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		vmName := r.URL.Query().Get("name")
		if !reflect.DeepEqual(vmName, newVMName) {
			t.Errorf("Request param `name` = %+v, expected %+v", vmName, newVMName)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Rename(1, "New-VM-Name")
	if err != nil {
		t.Errorf("VM.Rename returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Rename returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_Start(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/start", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Start(1)
	if err != nil {
		t.Errorf("VM.Start returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Start returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/reboot", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Reboot(1)
	if err != nil {
		t.Errorf("VM.Reboot returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Reboot returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_Reset(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/reset", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Reset(1)
	if err != nil {
		t.Errorf("VM.Reset returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Reset returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_PowerOff(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/poweroff", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.PowerOff(1)
	if err != nil {
		t.Errorf("VM.PowerOff returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.PowerOff returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_Shutdown(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/shutdown", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Shutdown(1)
	if err != nil {
		t.Errorf("VM.Shutdown returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Shutdown returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/delete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.Delete(1)
	if err != nil {
		t.Errorf("VM.Delete returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.Delete returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_ChangeCPU(t *testing.T) {
	setup()
	defer teardown()

	cpuCount := uint8(12)

	mux.HandleFunc("/v1/vm/1/cpu_change", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		newCPUCountStr := r.URL.Query().Get("cpu")
		newCPUCount, err := strconv.Atoi(newCPUCountStr)
		if err != nil {
			t.Errorf("VM.ChangeCPU returned error: %v", err)
		}

		if uint8(newCPUCount) != cpuCount {
			t.Errorf("Request param `name` = %+v, expected %+v", cpuCount, cpuCount)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.ChangeCPU(1, uint8(12))
	if err != nil {
		t.Errorf("VM.ChangeCPU returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.ChangeCPU returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_ChangeTariff(t *testing.T) {
	setup()
	defer teardown()

	tariffID := 1234

	mux.HandleFunc("/v1/vm/1/tariff_change", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		newTariffIDStr := r.URL.Query().Get("tariffId")
		newTariffID, err := strconv.Atoi(newTariffIDStr)
		if err != nil {
			t.Errorf("VM.ChangeTariff returned error: %v", err)
		}

		if tariffID != newTariffID {
			t.Errorf("Request param `tariffId` = %+v, expected %+v", newTariffID, tariffID)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.ChangeTariff(1, 1234)
	if err != nil {
		t.Errorf("VM.ChangeTariff returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.ChangeTariff returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_ChangePassword(t *testing.T) {
	setup()
	defer teardown()

	password := "newPassword"
	isSendPasswordTrue := "true"

	mux.HandleFunc("/v1/vm/1/password_change", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		newPassword := r.URL.Query().Get("password")
		if password != newPassword {
			t.Errorf("Request param `password` = %+v, expected %+v", newPassword, password)
		}

		isSendPassword := r.URL.Query().Get("sendPassword")
		if isSendPasswordTrue != isSendPassword {
			t.Errorf("Request param `sendPassword` = %+v, expected %+v", isSendPasswordTrue, isSendPassword)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.ChangePassword(1, password, true)
	if err != nil {
		t.Errorf("VM.ChangePassword returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.ChangePassword returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_ChangeMemory(t *testing.T) {
	setup()
	defer teardown()

	memory := uint(512)
	memoryStr := "512"
	isAllowRestartTrue := "true"

	mux.HandleFunc("/v1/vm/1/memory_change", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		newMemorySize := r.URL.Query().Get("memory")
		if memoryStr != newMemorySize {
			t.Errorf("Request param `memory` = %+v, expected %+v", newMemorySize, memory)
		}

		isAllowRestart := r.URL.Query().Get("allowRestart")
		if isAllowRestartTrue != isAllowRestart {
			t.Errorf("Request param `allowRestart` = %+v, expected %+v", isAllowRestartTrue, isAllowRestart)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.ChangeMemory(1, memory, true)
	if err != nil {
		t.Errorf("VM.ChangeMemory returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.ChangeMemory returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_ChangeDisk(t *testing.T) {
	setup()
	defer teardown()

	disk := uint(8192)
	diskStr := "8192"
	isAllowMemoryChangeTrue := "true"

	mux.HandleFunc("/v1/vm/1/disk_change", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		newDiskSize := r.URL.Query().Get("disk")
		if diskStr != newDiskSize {
			t.Errorf("Request param `disk` = %+v, expected %+v", newDiskSize, disk)
		}

		isAllowMemoryChange := r.URL.Query().Get("allowMemoryChange")
		if isAllowMemoryChangeTrue != isAllowMemoryChange {
			t.Errorf("Request param `allowMemoryChange` = %+v, expected %+v", isAllowMemoryChangeTrue, isAllowMemoryChange)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.ChangeDisk(1, disk, true)
	if err != nil {
		t.Errorf("VM.ChangeMemory returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.ChangeMemory returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_AddIP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/vm/1/ip_add", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.AddIP(1)
	if err != nil {
		t.Errorf("VM.AddIP returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.AddIP returned %+v, expected %+v", operationID, expected)
	}
}

func TestVMOperation_DeleteIP(t *testing.T) {
	setup()
	defer teardown()

	ip := "127.0.0.1"

	mux.HandleFunc("/v1/vm/1/ip_delete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		ipDeleted := r.URL.Query().Get("ip")
		if ip != ipDeleted {
			t.Errorf("Request param `ip` = %+v, expected %+v", ipDeleted, ip)
		}

		fmt.Fprintf(w, `{"status":"OK","operationId":1}`)
	})

	operationID, _, err := client.VM.DeleteIP(1, ip)
	if err != nil {
		t.Errorf("VM.DeleteIP returned error: %v", err)
	}

	expected := 1
	if !reflect.DeepEqual(*operationID, expected) {
		t.Errorf("VM.DeleteIP returned %+v, expected %+v", operationID, expected)
	}
}
