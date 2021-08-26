package data

import (
	"gorm.io/gorm"

	"github.com/yusrilmr/todolist/backend/models"
)

type LabelRepository struct {
	C *gorm.DB
}

// Create adds new Label
func (r *LabelRepository) Create(label *models.Label) error {
	err := r.C.Create(&label).Error
	return err
}

// Update modifies existing Label
func (r *LabelRepository) Update(label *models.Label) error {
	err := r.C.Model(&label).Updates(models.Label{
		Description: label.Description,
		Task: label.Task,
	}).Error
	return err
}

// Delete soft removes Label by id
func (r *LabelRepository) Delete(id string) error {
	err := r.C.Delete(&models.Label{}, id).Error
	return err
}
//func (r *LabelRepository) GetByTask(id string) []models.Label {
//	var labels []models.Label
//	taskid := bson.ObjectIdHex(id)
//	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
//	result := models.Label{}
//	for iter.Next(&result) {
//		labels = append(labels, result)
//	}
//	return labels
//}

// GetAll fetches all Labels
func (r *LabelRepository) GetAll() []models.Label {
	var labels []models.Label
	r.C.Find(&labels)
	return labels
}

// GetById fetches Label by id
func (r *LabelRepository) GetById(id string) (label models.Label, err error) {
	err = r.C.Preload("Task").First(&label, id).Error
	return
}