package ppeligible

import "myapp/model"

type PpSubItemRepositorer interface {
	FindAllPpSubItem() ([]model.PpSubItem, error)
	// FindByPpSubItem() ([]model.PpSubItem, error)
	// SavePpSubItem() ([]model.PpSubItem, error)
	// UpdatePpSubItem() ([]model.PpSubItem, error)
	// DeletePpSubItem() ([]model.PpSubItem, error)
}

type PpEligibleService struct {
	ppSubItemRepository PpSubItemRepositorer
}

func NewPpEligibleService(ppSubItemRepository PpSubItemRepositorer) *PpEligibleService {
	return &PpEligibleService{ppSubItemRepository}
}

func (s *PpEligibleService) GetPpEligible() ([]model.PpSubItem, error) {
	
	ppSubItems, err := s.ppSubItemRepository.FindAllPpSubItem()
	if err != nil {
		return nil, err
	}
	return ppSubItems, nil
}

// func (s *PpEligibleService) GetByPpEligible( ) ([]model.PpSubItem, error) {

// 	ppSubItems, err := s.ppSubItemRepository.FindByPpSubItem()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }

// func (s *PpEligibleService) SavePpEligible() ([]model.PpSubItem, error) {
// 	ppSubItems, err := s.ppSubItemRepository.SavePpSubItem()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }

// func (s *PpEligibleService) UpdatePpEligible() ([]model.PpSubItem, error) {
// 	ppSubItems, err := s.ppSubItemRepository.UpdatePpSubItem()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }

// func (s *PpEligibleService) DeletePpEligible() ([]model.PpSubItem, error) {
// 	ppSubItems, err := s.ppSubItemRepository.DeletePpSubItem()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ppSubItems, nil
// }
