package models

func ToPlane(msg *ADSBMessage) Plane {
	return Plane{
		ID:        msg.ICAO,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
		Altitude:  msg.Altitude,
		Heading:   msg.Heading,
		Speed:     msg.Speed,
		Timestamp: msg.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00"), // ISO 8601
		Type:	   msg.MessageType,
	}
}