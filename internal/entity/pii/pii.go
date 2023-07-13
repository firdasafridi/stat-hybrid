package pii

type TrxPII struct {
	IDCard      string `json:"id_card" gorm:"column:id_card"`
	FullName    string `json:"full_name" gorm:"column:full_name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
}

type TrxPIIHybridEncrypt struct {
	Data string `json:"wrap_data"`
}
