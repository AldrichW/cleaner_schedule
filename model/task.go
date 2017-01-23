package model

import (
    "aldrich_projects/cleaner_schedule/model"
)

type TaskInterface interface {
    AddRoommates(roommates []model.Roommate) bool
    GetAllRoommates() []model.Roommate
    GetAssignedRoommate() model.Roommate
    GetName() string
    GetDescription() string
    SetName(name string)
    SetDescription(desc string)
}

type Task struct {
    name, description string
    assignedRoommate model.Roommate
    possibleRoommates []model.Roommate
}

func (t *Task) AddRoommates(roommates []model.Roommate) bool {
    if(roommates == nil){
        return false
    }

    if t.possibleRoommates == nil {
        t.possibleRoommates = roommates
    } else {
        t.possibleRoommates = append(t.possibleRoommates, roommates)
    }

    return true
}

func (t *Task) GetAllRoommates() []model.Roommate {
    return t.possibleRoommates
} 

func (t *Task) GetAssignedRoommate() model.Roommate {
    return t.assignedRoommate
}

func (t *Task) SetName(name string) {
    t.name = name
}

func (t *Task) SetDescription(desc string) {
    t.description = desc
}