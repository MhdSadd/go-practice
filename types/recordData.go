package types


type RecordData struct {
	Title string
	Artist string
	Image  string
}

func (RecordData) TableName() string {
	return "vinylrecords"
}