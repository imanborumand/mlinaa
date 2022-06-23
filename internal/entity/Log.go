package entity

import "time"

type Log struct {
	ServiceName string      `json:"service_name"       example:"customer_mlinaa"`
	Request     interface{} `json:"request"  example:"{name: Mlinaa}"`
	UniqueId    string      `json:"unique_id"     example:"customer_mlinaa_36"`
	RequestTime time.Time   `json:"request_time"  example:"369582753"`
}
