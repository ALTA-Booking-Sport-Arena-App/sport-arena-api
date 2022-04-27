package helper

import "time"

type VenueRequestFormat struct {
	FacilityID []uint    `json:"facility_id" form:"facility_id"`
	Day        []string  `json:"day" form:"day"`
	VenueID    uint      `json:"venue_id" form:"venue_id"`
	OpenHour   time.Time `json:"open_hour" form:"open_hour"`
	CloseHour  time.Time `json:"close_hour" form:"close_hour"`
	Price      uint      `json:"price" form:"price"`
}
