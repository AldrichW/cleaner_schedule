package storage

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "database/sql"

    "aldrich_projects/cleaner_schedule/model"
)

type StorageServiceInterface interface {
    GetAllTasks() ([]model.Task, error)
    GetTask(name string) (model.Task, error)
    GetAllRoommates() ([]model.Roommate, error)
    GetRoommate(name string) (model.Roommate, error)
    AddRoommate(roommateName string) error
    AddTask(taskName string) error
    DeleteRoommate(rootName string) error
    DeleteTask(taskName string) error
    DeleteAllRoommates() error
    DeleteAllTasks() error
}

type StorageService struct {
    DSN string 
    DBSession *sql.DB
}

func NewStorageService() (*StorageService, error) {
    dsn := "aldrichw:password@/dbname"
    db, err := sql.Open("mysql", dsn)

    if err!= nil {
        return nil, fmt.Errorf("Failed to Create Storage Service:%s", err)
    }

    return &StorageService{
        DSN: dsn,
        DBSession: db,
    }, nil
}

func (s *StorageService) GetAllTasks() ([]model.Task, error) {

}

func (s *StorageService) GetTask(name string) (model.Task, error) {

}
 
func (s *StorageService) GetAllRoommates() ([]model.Roommate, error) {

}

func (s *StorageService) GetRoommate(name string) (model.Roommate, error) {

}

func (s *StorageService)  AddRoommate(roommateName string) error {

}

func (s *StorageService) AddTask(taskName string) error {

}

func (s *StorageService) DeleteRoommate(rootName string) error {

}

func (s* StorageService) DeleteTask(taskName string) error {

}

func (s *StorageService) DeleteAllRoommates() error {

}

func (s *StorageService) DeleteAllTasks() error {

}


