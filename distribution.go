package flops

const distributionBasePath = "v1/distribution"

// DistributionService is an interface for interfacing with the distribution
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type DistributionService interface {
	List() ([]Distribution, *Response, error)
}

// DistributionServiceOp handles communication with the image related methods of the
// Flops API.
type DistributionServiceOp struct {
	client *Client
}

var _ DistributionService = &DistributionServiceOp{}

// Distribution represents a Flops distribution
type Distribution struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Bitness     uint16 `json:"bitness"`
}

type distributionRoot struct {
	Status       string         `json:"status"`
	Distribution []Distribution `json:"result"`
}

// List lists all the distribution available.
func (s *DistributionServiceOp) List() ([]Distribution, *Response, error) {
	return s.list()
}

// Helper method for listing distribution
func (s *DistributionServiceOp) list() ([]Distribution, *Response, error) {
	req, err := s.client.NewRequest("GET", distributionBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(distributionRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Distribution, resp, err
}
