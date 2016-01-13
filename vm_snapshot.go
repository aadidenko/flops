package flops

import "fmt"

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
func (s *VMServiceOp) Snapshots(vmID int) ([]Snapshot, *Response, error) {
	return s.snapshotList(vmID)
}

// Helper method for listing snapshots
func (s *VMServiceOp) snapshotList(vmID int) ([]Snapshot, *Response, error) {
	path := fmt.Sprintf("%s/%d/snapshots", vmBasePath, vmID)
	req, err := s.client.NewRequest("GET", path, nil)
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
