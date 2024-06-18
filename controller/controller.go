package controller

import (
	"inventory/repository"
	"inventory/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"inventory/model"
)

var todoService service.TodoService = service.NewTodoService(repository.NewTodoRepository())

func AddTodo(c *gin.Context) {
	var todo model.Todo

	// 1 从请求中取出数据，并处理错误
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2 将数据存入数据库，并处理错误
	if err := todoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3 返回响应
	c.JSON(http.StatusOK, todo)
}

func GetAllTodos(c *gin.Context) {
	// 1 查询表中所有数据 并处理错误
	todoList, err := todoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// 2 返回所有数据
	c.JSON(http.StatusOK, todoList)
}

func GetTodoByID(c *gin.Context) {
	// 获取id 并处理错误
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 根据id查找数据 并处理错误
	todo, err := todoService.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回查找到的数据
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoByID(c *gin.Context) {
	// 获取id 并处理错误
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 根据id查找数据 并处理错误
	todo, err := todoService.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新数据 并处理错误
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := todoService.UpdateTodoByID(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// 返回更新成功的数据
	c.JSON(http.StatusOK, todo)
}

func DeleteTodoByID(c *gin.Context) {
	// 获取id 并处理错误
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 根据id删除数据 并处理错误
	if err := todoService.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回删除成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
