package controller

import (
	"net/http"
	"strconv"

	"github.com/lemonoa/LemonOA-Go/model"
	"github.com/lemonoa/LemonOA-Go/service"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

// RegisterRoutes 注册路由
func (c *TodoController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/todos")
	{
		api.GET("", c.GetTodoList)
		api.POST("", c.CreateTodo)
		api.PUT("/:id", c.UpdateTodo)
		api.DELETE("/:id", c.DeleteTodo)
		api.PUT("/:id/complete", c.MarkAsCompleted)
		api.PUT("/:id/uncomplete", c.MarkAsUncompleted)
	}
}

// GetTodoList 获取待办事项列表
func (c *TodoController) GetTodoList(ctx *gin.Context) {
	// TODO: 从JWT中获取userID
	userID := uint(1)
	status, _ := strconv.Atoi(ctx.Query("status"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	todos, total, err := c.todoService.GetTodoList(userID, status, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  todos,
		"total": total,
	})
}

// CreateTodo 创建待办事项
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo model.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 从JWT中获取userID
	todo.UserID = uint(1)

	if err := c.todoService.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

// UpdateTodo 更新待办事项
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var todo model.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = uint(id)
	// TODO: 从JWT中获取userID
	todo.UserID = uint(1)

	if err := c.todoService.UpdateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo 删除待办事项
func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.todoService.DeleteTodo(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// MarkAsCompleted 标记待办事项为已完成
func (c *TodoController) MarkAsCompleted(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.todoService.MarkAsCompleted(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// MarkAsUncompleted 标记待办事项为未完成
func (c *TodoController) MarkAsUncompleted(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// TODO: 从JWT中获取userID
	userID := uint(1)

	if err := c.todoService.MarkAsUncompleted(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
