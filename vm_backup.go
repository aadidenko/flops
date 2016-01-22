package flops

import "fmt"

// Backup represents a Flops backups of virtual machine
type Backup struct {
	ID        int
	Size      uint
	TimeAdded Timestamp
}

type BackupRequest struct{}

type backupsRoot struct {
	Status  string   `json:"status"`
	Backups []Backup `json:"result"`
}

// Backups lists all the backups available.
func (s *VMServiceOp) Backups(vmID int) ([]Backup, *Response, error) {
	return s.backupList(vmID)
}

// Helper method for listing backups
func (s *VMServiceOp) backupList(vmID int) ([]Backup, *Response, error) {
	path := fmt.Sprintf("%s/%d/backups", vmBasePath, vmID)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(backupsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Backups, resp, err
}
