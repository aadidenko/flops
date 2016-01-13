package flops

import "fmt"

const vmBasePath = "v1/vm"

// VMService is an interface for interfacing with the virtual machines
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type VMService interface {
	List() ([]VM, *Response, error)
	Get(int) (*VM, *Response, error)
	Rename(int, string) (*int, *Response, error)
	Start(int) (*int, *Response, error)
	Reboot(int) (*int, *Response, error)
	Reset(int) (*int, *Response, error)
	PowerOff(int) (*int, *Response, error)
	Shutdown(int) (*int, *Response, error)
	Delete(int) (*int, *Response, error)
	ChangeCPU(int, uint8) (*int, *Response, error)
	ChangeTariff(int, int) (*int, *Response, error)
	AddIP(int) (*int, *Response, error)
	DeleteIP(int, string) (*int, *Response, error)
	Snapshots(int) ([]Snapshot, *Response, error)
}

// VMServiceOp handles communication with the image related methods of the
// Flops API.
type VMServiceOp struct {
	client *Client
}

var _ VMService = &VMServiceOp{}

// VM represents a Flops virtual machine
type VM struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	InternalID       string        `json:"internalId"`
	Memory           uint          `json:"memory"`
	Disk             uint          `json:"disk"`
	CPU              uint8         `json:"cpu"`
	Bandwidth        uint          `json:"bandwidth"`
	TariffID         int           `json:"tariffId"`
	IPAddresses      []string      `json:"ipAddresses"`
	PrivateIPAddress string        `json:"privateIPAdresses"`
	State            string        `json:"state"`
	TimeAdded        Timestamp     `json:"timeAdded"`
	CurrentSnapshot  *int          `json:"currentSnapshot"`
	BackupPolicy     BackupPolicy  `json:"backupPolicy"`
	Distribution     Distribution  `json:"distribution"`
	OwnerUser        *string       `json:"ownerUser"`
	AccessUsers      []string      `json:"accessUsers"`
	PublicKeys       []VMPublicKey `json:"publicKeys"`
}

// VMPublicKey represents a Flops VM's public key
type VMPublicKey struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	PublicKey *string    `json:"publicKey"`
	TimeAdded *Timestamp `json:"timeAdded"`
	Installed bool       `json:"installed"`
	OwnerUser *string    `json:"ownerUser"`
}

type vmsRoot struct {
	Status string `json:"status"`
	VM     []VM   `json:"result"`
}

type vmRoot struct {
	Status string `json:"status"`
	VM     *VM    `json:"result"`
}

// List lists all the virtual machines available.
func (s *VMServiceOp) List() ([]VM, *Response, error) {
	return s.list()
}

// Get individual virtual machine
func (s *VMServiceOp) Get(vmID int) (*VM, *Response, error) {
	if vmID < 1 {
		return nil, nil, NewArgError("vmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d", vmBasePath, vmID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(vmRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.VM, resp, err
}

// Helper method for listing virtual machines
func (s *VMServiceOp) list() ([]VM, *Response, error) {
	req, err := s.client.NewRequest("GET", vmBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(vmsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.VM, resp, err
}
