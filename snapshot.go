package flops

const snapshotsBasePath = "v1/snapshots"

// SnapshotsService is an interface for interfacing with the snapshots
// endpoints of the Flops API
// See: http://support.flops.ru/index.php?/Knowledgebase/Article/View/24
type SnapshotsService interface {
	List() ([]Snapshot, *Response, error)
}

// SnapshotsServiceOp handles communication with the image related methods of the
// Flops API.
type SnapshotsServiceOp struct {
	client *Client
}

var _ SnapshotsService = &SnapshotsServiceOp{}

// Snapshot represents a Flops virtual machine snapshot
type Snapshot struct {
	ID               int
	Name             string
	Description      string
	ParentSnapshotID *int
	TimeAdded        Timestamp
}

type snapshotsRoot struct {
	Status    string     `json:"status"`
	Snapshots []Snapshot `json:"result"`
}

// List lists all the snapshots available.
func (s *SnapshotsServiceOp) List() ([]Snapshot, *Response, error) {
	return s.list()
}

// Helper method for listing snapshots
func (s *SnapshotsServiceOp) list() ([]Snapshot, *Response, error) {
	req, err := s.client.NewRequest("GET", snapshotsBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(snapshotsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Snapshots, resp, err
}
