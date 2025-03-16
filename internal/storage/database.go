package storage

import (
	"My_first_go_project/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

func Init(path string) *Database {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Cannot migrate database", err)
	}
	return &Database{DB: db}
}

func (d *Database) AddTask(task *models.Task) error {
	return d.DB.Create(&task).Error
}

func (d *Database) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := d.DB.Find(&tasks).Error
	return tasks, err
}
