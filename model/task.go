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
    Name, Description string
    AssignedRoommate model.Roommate // the roommate currently assigned to this Task
    PossibleRoommates []model.Roommate  // all possible roommates who can get selected for this Task
}

func (t *Task) AddRoommates(roommates []model.Roommate) bool {
    if roommates == nil {
        return false
    }

    if t.PossibleRoommates == nil {
        t.PossibleRoommates = roommates
    } else {
        t.PossibleRoommates = append(t.possibleRoommates, roommates)
    }

    return true
}

func (t *Task) GetAllRoommates() []model.Roommate {
    return t.PossibleRoommates
} 

func (t *Task) GetAssignedRoommate() model.Roommate {
    return t.AssignedRoommate
}

func (t *Task) SetName(name string) {
    t.Name = name
}

func (t *Task) SetDescription(desc string) {
    t.Description = desc
}