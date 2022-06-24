package entity

import "time"

type Log struct {
	RequestDate interface{} `json:"request_date"  example:"{name: Mlinaa}"`
	UniqueId    string      `json:"unique_id"     example:"customer_mlinaa_36"`
	RequestTime time.Time   `json:"request_time"  example:"369582753"`
}
