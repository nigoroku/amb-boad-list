package service

import (
	"fmt"
	"sort"
	"time"

	// "github.com/kzpolicy/user/models"
	"context"

	"github.com/volatiletech/null"
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
			mc.color_code,
			iaa.input_achievement_action_id,
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

	actionMap := make(map[null.Int]string)
	categoryMap := make(map[null.Int]string)
	lines := make(map[int]*models.Timeline)
	for _, tl := range timeline {

		if val, ok := lines[tl.AchievementID]; ok {

			if _, ok := actionMap[tl.InputAchievementActionID]; !ok {
				// 同じIDのアクションタイプは読み飛ばす
				if tl.ActionType.String != "" {
					ac := val.ActionTypes
					ac = append(ac, tl.ActionType.String)
					val.ActionTypes = ac

					actionMap[tl.InputAchievementActionID] = tl.ActionType.String
					lines[tl.AchievementID] = val
				}
			}

			if _, ok := categoryMap[tl.CategoryID]; !ok {
				// 同じIDのカテゴリは読み飛ばす

				if tl.CategoryID.Int != 0 {
					var ca generated.MCategory
					ca.CategoryID = tl.CategoryID.Int
					ca.Name = tl.CategoryName.String
					ca.ColorCode = tl.ColorCode

					cas := val.Categories
					cas = append(cas, ca)
					val.Categories = cas

					categoryMap[tl.CategoryID] = tl.CategoryName.String
					lines[tl.AchievementID] = val
				}
			}
		}

		if _, ok := lines[tl.AchievementID]; !ok {

			if tl.ReferenceURL.String != "" {
				// 参考URLからスクレイピングしてページ情報を取得する
				scrapingService := NewScrapingService()
				pageSummary := scrapingService.getPageSummary(tl.ReferenceURL.String)
				tl.InputPage = pageSummary
			}

			lines[tl.AchievementID] = tl
		}
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
			mc.color_code,
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

	actionMap := make(map[null.Int]string)
	categoryMap := make(map[null.Int]string)
	lines := make(map[int]*models.Timeline)
	for _, tl := range timeline {

		if val, ok := lines[tl.AchievementID]; ok {

			if _, ok := actionMap[tl.OutputAchievementActionID]; !ok {
				// 同じIDのアクションタイプは読み飛ばす
				if tl.ActionType.String != "" {
					ac := val.ActionTypes
					ac = append(ac, tl.ActionType.String)
					val.ActionTypes = ac

					actionMap[tl.OutputAchievementActionID] = tl.ActionType.String
					lines[tl.AchievementID] = val
				}
			}

			if _, ok := categoryMap[tl.CategoryID]; !ok {
				// 同じIDのカテゴリは読み飛ばす

				if tl.CategoryID.Int != 0 {
					var ca generated.MCategory
					ca.CategoryID = tl.CategoryID.Int
					ca.Name = tl.CategoryName.String
					ca.ColorCode = tl.ColorCode

					cas := val.Categories
					cas = append(cas, ca)
					val.Categories = cas

					categoryMap[tl.CategoryID] = tl.CategoryName.String
					lines[tl.AchievementID] = val
				}
			}
		}

		if _, ok := lines[tl.AchievementID]; !ok {

			if tl.ReferenceURL.String != "" {
				// 参考URLからスクレイピングしてページ情報を取得する
				scrapingService := NewScrapingService()
				pageSummary := scrapingService.getPageSummary(tl.ReferenceURL.String)
				tl.OutputPage = pageSummary
			}

			lines[tl.AchievementID] = tl
		}
	}

	return values(lines), err
}

// FindInputSelectedActions ユーザーに紐づくインプット実績のアクションを取得する
func (tl *TimelineService) FindInputSelectedActions(UserID int) ([]*generated.InputAchievementAction, error) {
	return generated.InputAchievementActions(qm.Where("user_id=?", UserID)).All(tl.ctx, tl.db)
}

// FindOutputSelectedActions ユーザーに紐づくアウトプット実績のアクションを取得する
func (tl *TimelineService) FindOutputSelectedActions(UserID int) ([]*generated.OutputAchievementAction, error) {
	return generated.OutputAchievementActions(qm.Where("user_id=?", UserID)).All(tl.ctx, tl.db)
}

// UpdateInputAction ユーザーに紐づくインプット実績を更新する
func (tl *TimelineService) UpdateInputAction(UserID int, ActionType string, Insert bool, AchievementID int) error {

	if Insert {
		a := &generated.InputAchievementAction{}
		a.InputAchievementID = AchievementID
		a.UserID = UserID
		a.ActionType = ActionType
		a.CreatedBy = UserID
		a.CreatedAt = time.Now()
		return a.Insert(tl.ctx, tl.db, boil.Infer())
	}

	_, err := generated.InputAchievementActions(qm.Where("input_achievement_id=? and  user_id=? and  action_type=?", AchievementID, UserID, ActionType)).DeleteAll(tl.ctx, tl.db)
	return err
}

// UpdateOutputAction ユーザーに紐づくアウトプット実績を更新する
func (tl *TimelineService) UpdateOutputAction(UserID int, ActionType string, Insert bool, AchievementID int) error {

	if Insert {
		a := &generated.OutputAchievementAction{}
		a.OutputAchievementID = AchievementID
		a.UserID = UserID
		a.ActionType = ActionType
		a.CreatedBy = UserID
		a.CreatedAt = time.Now()
		return a.Insert(tl.ctx, tl.db, boil.Infer())
	}

	_, err := generated.OutputAchievementActions(qm.Where("output_achievement_id=? and  user_id=? and action_type=?", AchievementID, UserID, ActionType)).DeleteAll(tl.ctx, tl.db)
	return err
}

func values(m map[int]*models.Timeline) []*models.Timeline {
	line := models.TimelineSlice{}
	for _, v := range m {
		line = append(line, v)
	}
	sort.Sort(line)
	return line
}
