package flops

const vmBasePath = "v1/vm"

// VMService is an interface for interfacing with the virtual machines
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type VMService interface {
	List() ([]VM, *Response, error)
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

type vmRoot struct {
	Status string `json:"status"`
	VM     []VM   `json:"result"`
}

// List lists all the virtual machines available.
func (s *VMServiceOp) List() ([]VM, *Response, error) {
	return s.list()
}

// Helper method for listing virtual machines
func (s *VMServiceOp) list() ([]VM, *Response, error) {
	req, err := s.client.NewRequest("GET", vmBasePath, nil)
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
