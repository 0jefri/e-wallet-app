package delivery

import (
	"e-wallet-app/config"
	"fmt"

	// "go-novel-api/delivery/controller"
	// "go-novel-api/repository"
	// "go-novel-api/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	// nvlUseCase usecase.NovelUsecase
	engine *gin.Engine
	host   string
}

func (a *appServer) Run() {
	// a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// dbConn, _ := config.NewDbConnection(cfg)
	dbConn, err := config.NewDbConnection(cfg)
	if err != nil {
		log.Fatalln("Failed to initialize database connection:", err)
		fmt.Println("database gagal")
	}

	if dbConn.Conn() == nil {
		log.Fatalln("Database connection is nil")
		fmt.Println("database berhasil")
	}
	// nvlRepo := repository.NewNovelRepository(dbConn.Conn())
	// nvlUseCase := usecase.NewNovelUsecase(nvlRepo)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)

	return &appServer{
		engine: engine,
		// nvlUseCase: nvlUseCase,
		host: host,
	}
}
