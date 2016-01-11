package flops

const pubKeysBasePath = "v1/pubkeys"

// PubKeysService is an interface for interfacing with the public keys
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type PubKeysService interface {
	List() ([]PubKey, *Response, error)
}

// PubKeysServiceOp handles communication with the image related methods of the
// Flops API.
type PubKeysServiceOp struct {
	client *Client
}

var _ PubKeysService = &PubKeysServiceOp{}

// PubKey represents a Flops public key
type PubKey struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	PublicKey string    `json:"publicKey"`
	TimeAdded Timestamp `json:"timeAdded"`
}

type pubKeysRoot struct {
	Status  string   `json:"status"`
	PubKeys []PubKey `json:"result"`
}

// List lists all the public keys available.
func (s *PubKeysServiceOp) List() ([]PubKey, *Response, error) {
	return s.list()
}

// Helper method for listing public keys
func (s *PubKeysServiceOp) list() ([]PubKey, *Response, error) {
	req, err := s.client.NewRequest("GET", pubKeysBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(pubKeysRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.PubKeys, resp, err
}
