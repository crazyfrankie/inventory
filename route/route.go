package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"inventory/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 设置 HTML 模板和静态文件目录
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 定义路由组
	v1Group := router.Group("/v1")
	{
		v1Group.POST("/todo", controller.AddTodo)
		v1Group.GET("/todo", controller.GetAllTodos)
		v1Group.GET("/todo/:id", controller.GetTodoByID)
		v1Group.PUT("/todo/:id", controller.UpdateTodoByID)
		v1Group.DELETE("/todo/:id", controller.DeleteTodoByID)
	}

	return router
}
