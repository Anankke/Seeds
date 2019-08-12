package main

import (
	"Seeds/src/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := router.Load()
	e.Run()
}
