package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vladqstrn/tasker-back/internal/config"
	"github.com/vladqstrn/tasker-back/internal/database"
	"github.com/vladqstrn/tasker-back/internal/server"
	"github.com/vladqstrn/tasker-back/internal/tasker/delivery/rest/handler"
	"github.com/vladqstrn/tasker-back/internal/tasker/repository"
)

func main() {
	//настройка конфига
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	if err := config.InitConfig(); err != nil {
		panic("failed to initialize config")
	}

	//инициализация роутера
	r := gin.Default()
	//инициализация БД
	db := database.InitDB()

	//инициализация слоя repository
	repo := repository.NewPostgresTaskRepository(db)

	//инициализация слоя delivery
	del := handler.NewTaskHandler(repo)

	//вызов конструктора приложения
	app := server.CreateApp(r, repo, del)

	//запуск приложения
	app.Run()
}
