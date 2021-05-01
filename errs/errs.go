package errs

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	
)

type Errs struct {
	HTTPStatusCode int    `json:"-"`
	Code           string `json:"code"`
	Message        string `joson:"massage"`
}

func New(httpStatusCode int, code, message string) *Errs {
	return &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Message:        message,
	}
}

func (e *Errs) Error() string {
	return fmt.Sprintf("code: %s, massage: %s", e.Code, e.Message)
}

func HTTPErrorHandler(err error, c echo.Context) {
	if _, ok := errors.Cause(err).(*Errs); ok {
		JSON(c, err)
		return
	}
	JSON(c ,New(http.StatusInternalServerError, "99999" , err.Error()))

}

func JSON(c echo.Context, err error) error {
	errs := errors.Cause(err).(*Errs)
	logrus.Errorf("ERROR: %v", errs)
	return c.JSON(errs.HTTPStatusCode, errs)


}
