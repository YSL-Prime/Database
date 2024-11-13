package entity

type Data struct {
	ID       uint    `json:"id"`
	Address  string  `json:"address" gorm:"not null" `
	Name     *string `json:"name"`
	Ds       string  `json:"ds" gorm:"size:500;not null"`
	Ok       *bool   `json:"ok"`
	Category *string `json:"category"`
	ImageUrl string  `json:"image_url" gorm:"not null"`
}
