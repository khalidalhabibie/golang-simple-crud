package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khalidalhabibie/depatu/model"
	"github.com/khalidalhabibie/depatu/repository"
)

type TaskHandler interface {
	AddTask(*gin.Context)
	GetAllTask(*gin.Context)
	GetTaskUserAdmin(*gin.Context)
	DeleteTask(*gin.Context)
}

type taskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler() TaskHandler {
	return &taskHandler{
		repo: repository.NewTaskRepository(),
	}
}

func (h *taskHandler) AddTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addtask, err := h.repo.AddTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, addtask)

}

func (h *taskHandler) GetAllTask(ctx *gin.Context) {
	task, err := h.repo.GetAllTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ctx.Writer.Status(), "error": err.Error()})
		return

	}
	codeStatus := ctx.Writer.Status()
	ctx.JSON(http.StatusOK, gin.H{"Data": task, "code": codeStatus})

}

func (h *taskHandler) GetTaskUserAdmin(ctx *gin.Context) {
	var task model.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		
	}
	fmt.Println(task.Assignedto)
	tasks, err := h.repo.GetTaskByUser(task.Assignedto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ctx.Writer.Status(), "error": err.Error()})
		return

	}
	codeStatus := ctx.Writer.Status()
	ctx.JSON(http.StatusOK, gin.H{"Data": tasks, "code": codeStatus})

}


func (h *taskHandler) DeleteTask(ctx *gin.Context) {
	var task model.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	

	fmt.Println(task)
	data, err := h.repo.DeleteTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	
	ctx.JSON(http.StatusOK, data)

}