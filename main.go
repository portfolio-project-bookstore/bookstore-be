package main

import (
	_config "bookstore/config"
	_ac "bookstore/delivery/controller/auth"
	_bc "bookstore/delivery/controller/book"
	_uc "bookstore/delivery/controller/user"
	_middleware "bookstore/delivery/middleware"
	_route "bookstore/delivery/route"
	_ar "bookstore/repository/auth"
	_br "bookstore/repository/book"
	_ur "bookstore/repository/user"
	_utility "bookstore/utility"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := _config.GetConfig()
	db := _utility.InitDB(config)
	user_repo := _ur.New(db)
	user_controller := _uc.New(user_repo)
	auth_repo := _ar.New(db)
	auth_controller := _ac.New(auth_repo)
	book_repo := _br.New(db)
	book_controller := _bc.New(book_repo)

	e := echo.New()
	e.Pre(_middleware.CustomLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	_route.RegisterPath(e, user_controller, auth_controller, book_controller)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
