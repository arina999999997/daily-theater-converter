package elichika

type DailyTheater struct {
	Language       string `json:"language"`
	DailyTheaterId int32  `json:"daily_theater_id"`
	Year           int32  `json:"year"`
	Month          int32  `json:"month"`
	Day            int32  `json:"day"`
	Title          string `json:"title"`
	DetailText     string `json:"detail_text"`
}
