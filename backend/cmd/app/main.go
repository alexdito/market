package main

import (
	"application/internal/app/common"
	"application/internal/app/providers"
	"application/internal/app/src/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	application, err := common.GetApp()

	if err != nil {
		fmt.Println("Ошибка подключения к БД!")
	}

	r := gin.Default()

	navigation := entity.Navigation{}
	application.DataBase.First(&navigation)

	controllers := providers.BuildDI(application)

	providers.RouteBuild(r, controllers)

	//repositories := repository.NewRepository(application.DataBase)
	//transactionService := service.NewService(repositories)
	//handlers := handler.NewHandler(transactionService)
	//
	////Баланс
	//r.GET("/balance", func(c *gin.Context) {
	//	c.JSON(handlers.GetBalance(c))
	//})
	//
	//r.GET("/info", func(c *gin.Context) {
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"balance":   "TEST",
	//		"currency":  "TEST",
	//		"operation": "Получение текущего баланса пользователя",
	//	})
	//})

	err = r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
