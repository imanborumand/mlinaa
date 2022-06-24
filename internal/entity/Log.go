package entity

import "time"

type Log struct {
	Data        interface{} `json:"data"  example:"{name: Mlinaa}" bson:"data"`
	UniqueId    string      `json:"unique_id"     example:"customer_mlinaa_36" bson:"unique_id"`
	RequestTime time.Time   `json:"request_time"  example:"369582753" bson:"request_time"`
}

type StoreRequestBody struct {
	Data interface{} `json:"data"`
}
