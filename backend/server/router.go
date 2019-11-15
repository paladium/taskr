package server

import (
	"taskr/controllers"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func router() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("v1")
	{
		tasksGroup := v1.Group("tasks")
		{
			tasks := new(controllers.TaskController)
			tasksGroup.GET("", tasks.Tasks)
			tasksGroup.PUT("move", tasks.MoveTask)
			tasksGroup.POST("add", tasks.AddTask)
		}
	}
	return router
}
