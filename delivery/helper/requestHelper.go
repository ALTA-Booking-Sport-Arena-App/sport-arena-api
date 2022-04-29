package helper

type VenueRequestFormat struct {
	FacilityID []uint   `json:"facility_id" form:"facility_id"`
	Day        []string `json:"day" form:"day"`
	VenueID    uint     `json:"venue_id" form:"venue_id"`
	OpenHour   string   `json:"open_hour" form:"open_hour"`
	CloseHour  string   `json:"close_hour" form:"close_hour"`
	Price      uint     `json:"price" form:"price"`
}
