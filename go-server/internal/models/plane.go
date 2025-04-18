package models

type Plane struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  int     `json:"altitude"`
	Heading   float64 `json:"heading"`
	Speed     float64 `json:"speed"`
	Timestamp string  `json:"timestamp"` // ISO 8601 string
	Type string 	  `json:"type"`
}