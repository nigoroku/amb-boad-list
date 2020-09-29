package models

import (
	"local.packages/models/generated"
)

type BoadList struct {
	User       generated.User                   `json:"user"`
	Todos      generated.TodoDetailSlice        `json:"todos"`
	OutputList generated.OutputAchievementSlice `json:"output_list"`
	InputList  generated.InputAchievementSlice  `json:"input_list"`
}
