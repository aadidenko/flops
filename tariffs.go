package flops

const tarifBasePath = "v1/tariffs"

// TariffsService is an interface for interfacing with the tariffs
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type TariffsService interface {
	List() ([]Tariff, *Response, error)
}

// TariffsServiceOp handles communication with the image related methods of the
// Flops API.
type TariffsServiceOp struct {
	client *Client
}

var _ TariffsService = &TariffsServiceOp{}

// Tariff represents a Flops tariff
type Tariff struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Memory     uint   `json:"memory"`
	Disk       uint   `json:"disk"`
	CPU        uint8  `json:"cpu"`
	IPCount    uint8  `json:"ipCount"`
	OnDemand   bool   `json:"onDemand"`
	ForWindows bool   `json:"forWindows"`
}

type tariffsRoot struct {
	Status  string   `json:"status"`
	Tariffs []Tariff `json:"result"`
}

// List lists all the tariffs available.
func (s *TariffsServiceOp) List() ([]Tariff, *Response, error) {
	return s.list()
}

// Helper method for listing tariffs
func (s *TariffsServiceOp) list() ([]Tariff, *Response, error) {
	req, err := s.client.NewRequest("GET", tarifBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(tariffsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Tariffs, resp, err
}
