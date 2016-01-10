package flops

const softwareBasePath = "v1/software"

// SoftwareService is an interface for interfacing with the software
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type SoftwareService interface {
	List() ([]Software, *Response, error)
}

// SoftwareServiceOp handles communication with the image related methods of the
// Flops API.
type SoftwareServiceOp struct {
	client *Client
}

var _ SoftwareService = &SoftwareServiceOp{}

// Software represents a Flops software
type Software struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type softwareRoot struct {
	Status   string     `json:"status"`
	Software []Software `json:"result"`
}

// List lists all the software available.
func (s *SoftwareServiceOp) List() ([]Software, *Response, error) {
	return s.list()
}

// Helper method for listing software
func (s *SoftwareServiceOp) list() ([]Software, *Response, error) {
	req, err := s.client.NewRequest("GET", softwareBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(softwareRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Software, resp, err
}
