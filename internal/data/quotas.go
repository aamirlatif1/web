package data

import "time"

type QuotaStatus string

const (
	StatusOK      QuotaStatus = "OK"
	StatusWarning QuotaStatus = "WARNING"
	StatusEExceed QuotaStatus = "EXCEEDED"
)

type Quota struct {
	QuotaID      int64       `json:"quota_id"`
	ResourceType string      `json:"resource_type"`
	Scope        string      `json:"scope"`
	ScopeID      string      `json:"scope_id"`
	Limits       Limits      `json:"limits"`
	Usage        Usage       `json:"usage"`
	Status       QuotaStatus `json:"status"`
}
type Limits struct {
	HardLimit       int    `json:"hard_limit"`
	SoftLimit       int    `json:"soft_limit"`
	RefreshInterval string `json:"refresh_interval"`
}
type Usage struct {
	Current     int       `json:"current"`
	LastRefresh time.Time `json:"last_refresh,omitzero"`
	NextRefresh time.Time `json:"next_refresh,omitzero"`
}
