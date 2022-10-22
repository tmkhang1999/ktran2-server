package structs

import "time"

type Response struct {
	SystemTime time.Time `json:"system_time"`
	Status     uint      `json:"status"`
}
