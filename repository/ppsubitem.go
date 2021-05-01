package repository

import (
	"myapp/model"

	"github.com/jinzhu/gorm"
	_"github.com/labstack/echo/v4"

)

type PpSubItemRepository struct {
	db *gorm.DB
}

func NewPpSubItemRepository(db *gorm.DB) *PpSubItemRepository {
	return &PpSubItemRepository{db}
}

func (r *PpSubItemRepository) FindAllPpSubItem() ([]model.PpSubItem, error) {
	var ppSubItems []model.PpSubItem
	// select * from PP_SUB_ITEM
	if err := r.db.Find(&ppSubItems).Error; err != nil {
		return nil, err
	}
	return ppSubItems, nil
}

// func (r *PpSubItemRepository) FindByPpSubItem() ([]model.PpSubItem, error) {
	
// 	//id := echo.Context.Param("id")
// 	var ppSubItems []model.PpSubItem
// 	// select ID from PP_SUB_ITEM
// 	if err := r.db.Find(&ppSubItems).Error; err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }

// func (r *PpSubItemRepository) SavePpSubItem() ([]model.PpSubItem, error) {

// 	var ppSubItems []model.PpSubItem

// 	if err := r.db.Save(&ppSubItems).Error; err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }

// func (r *PpSubItemRepository) UpdatePpSubItem() ([]model.PpSubItem, error) {

// 	//id := c.Param("id")
// 	var ppSubItems []model.PpSubItem

// 	if err := r.db.Find(&ppSubItems).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := r.db.Save(&ppSubItems).Error; err != nil {
// 		return nil, err
// 	}

// 	return ppSubItems, nil

// }

// func (r *PpSubItemRepository) DeletePpSubItem() ([]model.PpSubItem, error) {

// 	//id := c.Param("id")
// 	var ppSubItems []model.PpSubItem

// 	if err := r.db.Find(&ppSubItems).Error; err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }
