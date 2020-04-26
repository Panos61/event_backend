package models

// Events => Primary struct
type Events struct {
	Category    string
	Title       string
	Description string

	TimeType int
	Date     int

	Comments string
	AgeLimit bool
	Payment  bool
}
