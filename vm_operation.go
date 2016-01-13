package flops

import (
	"fmt"
	"strconv"
	"strings"
)

// Rename a virtual machine
func (s *VMServiceOp) Rename(vmID int, name string) (*int, *Response, error) {
	operation := "rename"
	opts := &operationOpts{"name": name}

	return s.doOperation(vmID, operation, opts)
}

// Start a virtual machine
func (s *VMServiceOp) Start(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "start", nil)
}

// Reboot a virtual machine
func (s *VMServiceOp) Reboot(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "reboot", nil)
}

// Reset a virtual machine
func (s *VMServiceOp) Reset(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "reset", nil)
}

// PowerOff a virtual machine
func (s *VMServiceOp) PowerOff(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "poweroff", nil)
}

// Shutdown a virtual machine
func (s *VMServiceOp) Shutdown(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "shutdown", nil)
}

// Delete a virtual machine
func (s *VMServiceOp) Delete(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "delete", nil)
}

// ChangeCPU changes CPU core counts for the virtual machine
func (s *VMServiceOp) ChangeCPU(vmID int, count uint8) (*int, *Response, error) {
	operation := "cpu_change"
	opts := &operationOpts{"cpu": strconv.Itoa(int(count))}
	return s.doOperation(vmID, operation, opts)
}

// ChangeTariff changes tariff for the virtual machine
func (s *VMServiceOp) ChangeTariff(vmID int, tariffID int) (*int, *Response, error) {
	operation := "tariff_change"
	opts := &operationOpts{"tariffId": strconv.Itoa(tariffID)}
	return s.doOperation(vmID, operation, opts)
}

// AddIP adds IP address for the virtual machine
func (s *VMServiceOp) AddIP(vmID int) (*int, *Response, error) {
	return s.doOperation(vmID, "ip_add", nil)
}

// DeleteIP removes IP address for the virtual machine
func (s *VMServiceOp) DeleteIP(vmID int, ip string) (*int, *Response, error) {
	operation := "ip_delete"
	opts := &operationOpts{"ip": ip}
	return s.doOperation(vmID, operation, opts)
}

func (s *VMServiceOp) doOperation(vmID int, op string, opts *operationOpts) (*int, *Response, error) {
	if vmID < 1 {
		return nil, nil, NewArgError("vmID", "can't be less than 1")
	}

	if op == "" {
		return nil, nil, NewArgError("op", "can't be blank")
	}

	path := fmt.Sprintf("%s/%d/%s", vmBasePath, vmID, op)
	if opts != nil {
		params := []string{}
		for name, value := range *opts {
			params = append(params, fmt.Sprintf("%s=%s", name, value))
		}
		path = fmt.Sprintf("%s?%s", path, strings.Join(params, "&"))
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	opResp := new(operationResponse)
	resp, err := s.client.Do(req, opResp)
	if err != nil {
		return nil, resp, err
	}

	return &opResp.OperationID, resp, err
}
