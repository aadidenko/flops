package flops

import "fmt"

const operationBasePath = "v1/operation"

// OperationService is an interface for interfacing with the operation
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type OperationService interface {
	Get(int) (*Operation, *Response, error)
}

// OperationServiceOp handles communication with the image related methods of the
// Flops API.
type OperationServiceOp struct {
	client *Client
}

var _ OperationService = &OperationServiceOp{}

// Operation represents a Flops operation
type Operation struct {
	ID           int     `json:"id"`
	VMID         int     `json:"vmId"`
	Status       string  `json:"status"`
	Type         string  `json:"operationType"`
	Percentage   uint    `json:"percentage"`
	ErrorMessage *string `json:"errorMessage,omitemty"`
	ErrorCode    *string `json:"errorMessage,omitemty"`
}

type operationOpts map[string]string

type operationRoot struct {
	Status    string     `json:"status"`
	Operation *Operation `json:"result"`
}

type operationResponse struct {
	Status      string `json:"status"`
	OperationID int    `json:"operationId"`
}

// Get operation.
func (s *OperationServiceOp) Get(opID int) (*Operation, *Response, error) {
	if opID < 1 {
		return nil, nil, NewArgError("opID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d", operationBasePath, opID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(operationRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Operation, resp, err
}
