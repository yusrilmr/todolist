package data

import (
	"gorm.io/gorm"

	"github.com/yusrilmr/todolist/backend/models"
)

type TaskRepository struct {
	C *gorm.DB
}

func (r *TaskRepository) Create(task *models.Task) error {
	task.Status = "Created"
	err := r.C.Create(&task).Error
	return err
}

func (r *TaskRepository) Update(task *models.Task) error {
	err := r.C.Model(&task).Updates(models.Task{
		Name: task.Name,
		Description: task.Description,
		Due: task.Due,
		Status: task.Status,
		UserID: task.UserID,
		Label: task.Label,
	}).Error
	return err
}

func (r *TaskRepository) Delete(id string) error {
	err := r.C.Delete(&models.Task{}, id).Error
	return err
}

func (r *TaskRepository) GetAll() []models.Task {
	//var tasks []models.Task
	//iter := r.C.Find(nil).Iter()
	//result := models.Task{}
	//for iter.Next(&result) {
	//	tasks = append(tasks, result)
	//}
	var tasks []models.Task
	//tasks = r.C.Find(&tasks)
	//results := r.C.Find(&tasks)
	return tasks
}

func (r *TaskRepository) GetById(id string) (task models.Task, err error) {
	err = r.C.Preload("Label").First(&task, id).Error
	return
}

func (r *TaskRepository) GetByUser(user string) []models.Task {
	//var tasks []models.Task
	//iter := r.C.Find(bson.M{"createdby": user}).Iter()
	//result := models.Task{}
	//for iter.Next(&result) {
	//	tasks = append(tasks, result)
	//}
	var tasks []models.Task
	return tasks
}