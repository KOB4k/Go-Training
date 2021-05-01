package main

import (
	"log"
	"net/http"
	"strings"

	"myapp/errs"
	customMdw "myapp/middleware"
	"myapp/model"
	"myapp/repository"
	"myapp/service/ppeligible"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	initViper()
}

func main() {
	// sample := viper.GetString("app.name")
	// fmt.Println(sample)
	var (
		e  = echo.New()
		db = newMariaConn()
	)

	//g := e.Group("/hello", middleware.Logger())

	e.HideBanner = true
	// e.HidePort = true
	e.HTTPErrorHandler = errs.HTTPErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(customMdw.AddHeader)
	e.Use(customMdw.LogRequestInfo())
	e.Use(customMdw.LogResponseInfo())

	g := e.Group("/hello")
	g.GET("/1/:text", GetHello)
	g.GET("/2", GetHelloQueryString)

	e.POST("/user", PostUser)

	ppEligibleHandler := ppeligible.NewPpEligibleHandler(
		ppeligible.NewPpEligibleService(
			repository.NewPpSubItemRepository(db),
		),
	)

	e.GET("/ppeligible", ppEligibleHandler.GetPpEligible) 
	// e.GET("ppeligible/:id", ppEligibleHandler.GetByPpEligible)
	// e.POST("/ppeligible", ppEligibleHandler.SavePpEligible)
	// e.PUT("/ppeligible/:id", ppEligibleHandler.UpdatePpEligible)
	// e.DELETE("/ppeligible/:id", ppEligibleHandler.DeletePpEligible)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetHello(c echo.Context) error {
	text := c.Param("text")
	return c.String(http.StatusOK, text)
}

func GetHelloQueryString(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	// id:1, name: leon
	return c.String(http.StatusOK, "id: "+id+",name: "+name)

}

func PostUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
	//return errors.New("exception error 500")
	//return errs.JSON(c, errs.New(http.StatusBadRequest, "40000" , "bad request error"))
}

func newMariaConn() *gorm.DB {

	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatalf("cannot open sqlite3 connection:%+v", err)
	}

	db.AutoMigrate(&model.PpSubItem{})

	// db.Create(&model.PpSubItem{
	// 	SubItemNo: "CO-19",
	// 	AppAlias: "KAI",
	// })

	return db
}

func initViper() {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
