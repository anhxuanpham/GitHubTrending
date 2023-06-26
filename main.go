package main

import (
	"GitHubTrending/db"
	"GitHubTrending/handler"
	"GitHubTrending/log"
	"GitHubTrending/repository/repo_impl"
	"GitHubTrending/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	fmt.Println("main function")
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5433,
		UserName: "postgres",
		Password: "postgres",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1234"))
}

//func logErr(errMsg string) {
//	_, file, line, _ := runtime.Caller(1)
//	log.WithFields(log.Fields{
//		"file": file,
//		"line": line,
//	}).Error(errMsg)
//
//}
