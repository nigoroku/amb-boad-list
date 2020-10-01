package service

import (
	"fmt"

	// "github.com/kzpolicy/user/models"
	"context"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models"
	"local.packages/models/generated"
)

type BoadService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewBoadService() *BoadService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &BoadService{ctx, db}
}

func (b *BoadService) FindBoadList() ([]*models.BoadList, error) {

	users, err := generated.Users(qm.Load(qm.Rels("Todos", "TodoDetails"), qm.Where("created_at = (select max(created_at) from todos)")),
		qm.Load("InputAchievements", qm.Where("created_at = (select max(created_at) from input_achievements)")),
		qm.Load("OutputAchievements", qm.Where("created_at = (select max(created_at) from output_achievements)"))).All(b.ctx, b.db)

	var boadList []*models.BoadList
	for _, u := range users {
		boad := &models.BoadList{}
		r := u.R
		t := r.Todos
		user := generated.User{}
		user.UserID = u.UserID
		user.AccountName = u.AccountName
		user.AccountImg = u.AccountImg
		user.Introduction = u.Introduction
		user.ContentType = u.ContentType
		boad.User = user
		if len(t) > 0 {
			boad.Todos = t[0].R.TodoDetails
		}
		if len(r.InputAchievements) > 0 {
			input := generated.InputAchievement{}
			in := r.InputAchievements[0]
			input.ReferenceURL = in.ReferenceURL
			input.Summary = in.Summary
			input.InputTime = in.InputTime
			boad.Input = input
			scrapingService := NewScrapingService()
			if in.ReferenceURL.String != "" {
				pageSummary := scrapingService.getPageSummary(in.ReferenceURL.String)
				boad.InputPage = pageSummary
			}
		}
		if len(r.OutputAchievements) > 0 {
			output := generated.OutputAchievement{}
			ou := r.OutputAchievements[0]
			output.ReferenceURL = ou.ReferenceURL
			output.Summary = ou.Summary
			output.OutputTime = ou.OutputTime
			boad.Output = output
			scrapingService := NewScrapingService()
			if ou.ReferenceURL.String != "" {
				pageSummary := scrapingService.getPageSummary(ou.ReferenceURL.String)
				boad.OutputPage = pageSummary
			}
		}
		boadList = append(boadList, boad)
	}

	if err != nil {
		fmt.Println("error %v", err)
		return boadList, err
	}
	return boadList, err
}
