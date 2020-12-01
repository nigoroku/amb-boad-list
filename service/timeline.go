package service

import (
	"fmt"

	// "github.com/kzpolicy/user/models"
	"context"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/generated"
	"local.packages/models"
)

type TimelineService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewTimelineService() *TimelineService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &TimelineService{ctx, db}
}

func (tl *TimelineService) FindInputTimeline() ([]models.Timeline, error) {

	var timeline []models.Timeline
	var query = `
			select
			u.user_id,
			u.account_name,
			u.account_img,
			u.introduction,
			u.content_type,
			ia.input_achievement_id as achievement_id,
			ia.reference_url,
			ia.summary,
			ia.created_at,
			mc.category_id,
			mc.name as category_name,
			iaa.action_type
		from
			input_achievements ia
		left outer join users u ON
			ia.user_id = u.user_id
		left outer join input_achievement_tags iat ON
			ia.input_achievement_id = iat.input_achievement_id
		left outer join m_categories mc ON
			iat.category_id = mc.category_id
		left outer join input_achievement_actions iaa ON
			ia.input_achievement_id = iaa.input_achievement_id
		order by
			ia.created_at DESC`

	err := generated.NewQuery(qm.SQL(query)).Bind(tl.ctx, tl.db, &timeline)

	// var timelines []*models.Timeline
	// for _, u := range users {
	// 	timeline := &models.Timeline{}
	// 	r := u.R
	// 	user := generated.User{}
	// 	user.UserID = u.UserID
	// 	user.Email = u.Email
	// 	user.AccountName = u.AccountName
	// 	user.AccountImg = u.AccountImg
	// 	user.Introduction = u.Introduction
	// 	user.ContentType = u.ContentType
	// 	timeline.User = user
	// 	if len(r.InputAchievements) > 0 {
	// 		input := generated.InputAchievement{}
	// 		in := r.InputAchievements[0]
	// 		input.ReferenceURL = in.ReferenceURL
	// 		input.Summary = in.Summary
	// 		input.InputTime = in.InputTime
	// 		input.CreatedAt = in.CreatedAt
	// 		timeline.Input = input
	// 		scrapingService := NewScrapingService()
	// 		if in.ReferenceURL.String != "" {
	// 			pageSummary := scrapingService.getPageSummary(in.ReferenceURL.String)
	// 			timeline.InputPage = pageSummary
	// 		}
	// 	}
	// if len(r.OutputAchievements) > 0 {
	// 	output := generated.OutputAchievement{}
	// 	ou := r.OutputAchievements[0]
	// 	output.ReferenceURL = ou.ReferenceURL
	// 	output.Summary = ou.Summary
	// 	output.OutputTime = ou.OutputTime
	// 	output.CreatedAt = ou.CreatedAt
	// 	timeline.Output = output
	// 	scrapingService := NewScrapingService()
	// 	if ou.ReferenceURL.String != "" {
	// 		pageSummary := scrapingService.getPageSummary(ou.ReferenceURL.String)
	// 		timeline.OutputPage = pageSummary
	// 	}
	// }
	// 	timelines = append(timelines, timeline)
	// }

	if err != nil {
		fmt.Printf("error %v", err)
		return timeline, err
	}
	return timeline, err
}
