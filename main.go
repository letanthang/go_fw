package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/letanthang/go_fw/config"
	"github.com/letanthang/go_fw/db"
	"github.com/letanthang/go_fw/route"
)

func main() {
	fmt.Println(config.Config)
	fmt.Println(db.Client)
	e := echo.New()
	route.Public(e)

	err := e.Start(":9090")

	if err != nil {
		fmt.Println("Cannot start server", err)
	}

}
