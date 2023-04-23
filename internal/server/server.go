package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-back/internal/config"
	"github.com/vladqstrn/tasker-back/internal/mw"
	"github.com/vladqstrn/tasker-back/internal/tasker/delivery/rest/handler"
	"github.com/vladqstrn/tasker-back/internal/tasker/delivery/rest/routes"
	"github.com/vladqstrn/tasker-back/internal/tasker/repository"
)

type App struct {
	r    *gin.Engine
	repo *repository.PostgresTaskRepository
	del  *handler.TaskHandler
}

func CreateApp(r *gin.Engine, repo *repository.PostgresTaskRepository, del *handler.TaskHandler) *App {
	return &App{
		r:    r,
		repo: repo,
		del:  del,
	}
}

func (app *App) Run() {
	mw.CORSMiddleware(app.r)
	routes.TaskRoutes(app.r, *app.repo)
	app.r.Run(config.Domain + ":" + config.AppPort)
}
