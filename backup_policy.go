package flops

type BackupPolicy struct {
	Quantity  uint `json:"quantity"`
	Frequency uint `json:"frequency"`
}
