package models




type Product struct {

	Id int64 `gorm:"PrimaryKey" json:"Id"`
	Title string `gorm:"type:varchar(25)" json:"Title"`
	Outhor string `gorm:"type:varchar(25)" json:"Outhor"`
	Desci string `gorm:"type:varchar(25)" json:"Desci"`
//	Product string 'gorm:"PrimaryKey" json:"id"'


}