package models

type User struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
	Prodi string `json:"prodi"`
}

type UserReq struct {
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
	Prodi string `json:"prodi"`
}
