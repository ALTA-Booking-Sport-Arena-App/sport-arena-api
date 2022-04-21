package helper

type UpdateEventRequest struct {
	Name     string `json:"name" form:"name"`
	Date     string `json:"date" form:"date"`
	Location string `json:"location" form:"location"`
	Details  string `json:"details" form:"details"`
	Quota    int    `json:"quota" form:"quota"`
	// Image        string `json:"image" form:"image"`
}
