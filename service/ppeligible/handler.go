package ppeligible

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PpEligibleServicer interface {
	GetPpEligible() ([]model.PpSubItem, error)
	// GetByPpEligible() ([]model.PpSubItem, error)
	// SavePpSubItem() ([]model.PpSubItem, error)
	// UpdatePpEligible() ([]model.PpSubItem, error)
	// DeletePpEligible() ([]model.PpSubItem, error)
}

type PpEligibleHandler struct {
	ppEligibleService PpEligibleServicer
}

func NewPpEligibleHandler(ppEligibleService PpEligibleServicer) *PpEligibleHandler {
	return &PpEligibleHandler{ppEligibleService}
}

func (h *PpEligibleHandler) GetPpEligible(c echo.Context) error {

	res, err := h.ppEligibleService.GetPpEligible()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// func (h *PpEligibleHandler) GetByPpEligible(c echo.Context) error {

// 	//id,_:= strconv.Atoi(c.Param("id"))

// 	//id := echo.Context.Param("id")

// 	res, err := h.ppEligibleService.GetByPpEligible()

// 	if err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

// func (h *PpEligibleHandler) SavePpEligible(c echo.Context) error {

// 	res, err := h.ppEligibleService.SavePpSubItem()
// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

// func (h *PpEligibleHandler) UpdatePpEligible(c echo.Context) error {

// 	res, err := h.ppEligibleService.UpdatePpEligible()
// 	if err != nil {

// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

// func (h *PpEligibleHandler) DeletePpEligible(c echo.Context) error {

// 	res, err := h.ppEligibleService.DeletePpEligible()
// 	if err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	return c.JSON(http.StatusOK, res)
// }
