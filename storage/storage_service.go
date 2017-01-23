package storage

type StorageServiceInterface interface {
    NewStorageService()
    GetAllRoommates()
    AddRoommate(roommateName string)
    AddTask (taskName string)
}

type StorageService struct {
    
}

func (s *StorageService) NewStorageService() *StorageService {
    return &StorageService{}

}
