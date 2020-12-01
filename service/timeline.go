package service

import (
	"fmt"
	"sort"

	// "github.com/kzpolicy/user/models"
	"context"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models"
	"local.packages/models/generated"
)

// TimelineService タイムラインサービス
type TimelineService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

// NewTimelineService コンストラクタ
func NewTimelineService() *TimelineService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &TimelineService{ctx, db}
}

// FindInputTimeline input実績のタイムラインに表示するデータを取得する
func (tl *TimelineService) FindInputTimeline() ([]*models.Timeline, error) {

	timeline := models.TimelineSlice{}
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

	err := models.NewQuery(qm.SQL(query)).Bind(tl.ctx, tl.db, &timeline)

	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	lines := make(map[int]*models.Timeline)
	for _, tl := range timeline {

		if val, ok := lines[tl.AchievementID]; ok {
			// アクション設定
			if !val.Lgtm {
				val.Lgtm = tl.ActionType.String == "1"
			}
			if !val.Stock {
				val.Stock = tl.ActionType.String == "2"
			}
			// カテゴリ設定
			if tl.CategoryID.Int != 0 {
				var ca generated.MCategory
				ca.CategoryID = tl.CategoryID.Int
				ca.Name = tl.CategoryName.String

				cas := val.Categories
				cas = append(cas, ca)
			}
			lines[tl.AchievementID] = val
			continue
		}

		var ca generated.MCategory
		var cas []generated.MCategory
		if tl.CategoryID.Int != 0 {
			ca.CategoryID = tl.CategoryID.Int
			ca.Name = tl.CategoryName.String
		}
		cas = append(cas, ca)
		tl.Categories = cas

		if tl.ReferenceURL.String != "" {
			// 参考URLからスクレイピングしてページ情報を取得する
			scrapingService := NewScrapingService()
			pageSummary := scrapingService.getPageSummary(tl.ReferenceURL.String)
			tl.InputPage = pageSummary
		}

		lines[tl.AchievementID] = tl
	}

	return values(lines), err
}

// FindOutputTimeline output実績のタイムラインに表示するデータを取得する
func (tl *TimelineService) FindOutputTimeline() ([]*models.Timeline, error) {

	timeline := models.TimelineSlice{}
	var query = `
			select
			u.user_id,
			u.account_name,
			u.account_img,
			u.introduction,
			u.content_type,
			ia.output_achievement_id as achievement_id,
			ia.reference_url,
			ia.summary,
			ia.created_at,
			mc.category_id,
			mc.name as category_name,
			iaa.action_type
		from
			output_achievements ia
		left outer join users u ON
			ia.user_id = u.user_id
		left outer join output_achievement_tags iat ON
			ia.output_achievement_id = iat.output_achievement_id
		left outer join m_categories mc ON
			iat.category_id = mc.category_id
		left outer join output_achievement_actions iaa ON
			ia.output_achievement_id = iaa.output_achievement_id
		order by
			ia.created_at DESC`

	err := models.NewQuery(qm.SQL(query)).Bind(tl.ctx, tl.db, &timeline)

	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	lines := make(map[int]*models.Timeline)
	for _, tl := range timeline {

		if val, ok := lines[tl.AchievementID]; ok {
			// アクション設定
			if !val.Lgtm {
				val.Lgtm = tl.ActionType.String == "1"
			}
			if !val.Stock {
				val.Stock = tl.ActionType.String == "2"
			}
			// カテゴリ設定
			if tl.CategoryID.Int != 0 {
				var ca generated.MCategory
				ca.CategoryID = tl.CategoryID.Int
				ca.Name = tl.CategoryName.String

				cas := val.Categories
				cas = append(cas, ca)
			}
			lines[tl.AchievementID] = val
			continue
		}

		var ca generated.MCategory
		var cas []generated.MCategory
		if tl.CategoryID.Int != 0 {
			ca.CategoryID = tl.CategoryID.Int
			ca.Name = tl.CategoryName.String
		}
		cas = append(cas, ca)
		tl.Categories = cas

		if tl.ReferenceURL.String != "" {
			// 参考URLからスクレイピングしてページ情報を取得する
			scrapingService := NewScrapingService()
			pageSummary := scrapingService.getPageSummary(tl.ReferenceURL.String)
			tl.OutputPage = pageSummary
		}

		lines[tl.AchievementID] = tl
	}

	return values(lines), err
}

func values(m map[int]*models.Timeline) []*models.Timeline {
	line := models.TimelineSlice{}
	for _, v := range m {
		line = append(line, v)
	}
	sort.Sort(line)
	return line
}
