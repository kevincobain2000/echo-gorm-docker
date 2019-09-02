package main

import (
	"os"
	"time"

	//"github.com/facebookgo/grace/gracehttp"
	"github.com/kevincobain2000/echo-gorm-docker/route"
	"github.com/tylerb/graceful"
)

func main() {
	// db.Init()
	e := route.Init()
	port := os.Args[1]

	//e.Logger.Fatal(e.Start(":" + port))
	e.Server.Addr = ":" + port
	graceful.ListenAndServe(e.Server, 5*time.Second)

	// e.Logger.Fatal(gracehttp.Serve(e.Server))
}
