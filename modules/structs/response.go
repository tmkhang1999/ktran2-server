package structs

import "time"

type Response struct {
	SystemTime time.Time `json:"system_time"`
	Status     uint      `json:"status"`
}

type HttpStatusMessage struct {
	Table       string `json:"table"`
	RecordCount *int64 `json:"record_count"`
}
