package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mnsh5/todo-api/services"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(s services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: s}
}

func (t *TaskHandler) FindAll(c echo.Context) error {
	log.Info("findAll tasks")
	return t.taskService.FindAll(c)
}

func (t *TaskHandler) FindById(c echo.Context) error {
	log.Info("Attempting to find task by ID")
	return t.taskService.FindById(c)
}

func (t *TaskHandler) Create(c echo.Context) error {
	log.Info("Attempting to create a new task")
	return t.taskService.Create(c)
}

func (t *TaskHandler) Update(c echo.Context) error {
	log.Info("Attempting to update a task")
	return t.taskService.Update(c)
}

func (t *TaskHandler) Delete(c echo.Context) error {
	log.Info("Attempting to update a task")
	return t.taskService.Delete(c)
}
