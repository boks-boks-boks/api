package model

import (
	uuid "github.com/google/uuid"
)

type Label struct {
	Id          uuid.UUID
	Title       string
	Color       string
	Description string
}

func InitLabel(title, color, description string) *Label {
	return &Label{Id: uuid.New(), Title: title, Color: color, Description: description}
}
