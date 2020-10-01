package models

import (
	"local.packages/models/generated"
)

type BoadList struct {
	User       generated.User              `json:"user"`
	Todos      generated.TodoDetailSlice   `json:"todos"`
	Output     generated.OutputAchievement `json:"output_list"`
	Input      generated.InputAchievement  `json:"input_list"`
	InputPage  PageSummary                 `json:"input_page_summary"`
	OutputPage PageSummary                 `json:"output_page_summary"`
}
