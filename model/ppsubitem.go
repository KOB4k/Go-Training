package model

type PpSubItem struct {
	// Id        uint   `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"column:FIRST_NAME" json:"firstName"`
	LastName  string `gorm:"column:SUB_ITEM_NO" json:"lastName"`
	SubItemNo string `gorm:"column:SUB_ITEM_NO" json:"subItemNo"`
	AppAlias  string `gorm:"column:APP_ALIAS" json:"appAlias"`
}

func (PpSubItem) TableName() string {
	return "PP_SUB_ITEM"
}
