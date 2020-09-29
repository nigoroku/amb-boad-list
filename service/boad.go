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
	// データ取得
	// var result = new(models.BoadListSliceModel)

	users, err := generated.Users(qm.Load(qm.Rels("Todos", "TodoDetails")), qm.Load("InputAchievements"), qm.Load("OutputAchievements")).All(b.ctx, b.db)

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
		boad.User = user
		if len(t) > 0 {
			boad.Todos = t[0].R.TodoDetails
		}
		if len(r.InputAchievements) > 0 {
			boad.InputList = r.InputAchievements
		}
		if len(r.OutputAchievements) > 0 {
			boad.OutputList = r.OutputAchievements
		}
		boadList = append(boadList, boad)
	}

	// var query = `select
	// 				u.user_id,
	// 				u.account_name,
	// 				u.account_img,
	// 				td.todo_id,
	// 				td.content
	// 			from
	// 				users u
	// 			left outer join todos t on
	// 				u.user_id = t.user_id
	// 			left outer join todo_details td on
	// 				t.todo_id = td.todo_id
	// 			left outer join output_achievements oa on
	// 				u.user_id = oa.user_id
	// 			left outer join input_achievements ia on
	// 				u.user_id = ia.user_id`
	// err := generated.NewQuery(qm.SQL(query)).Bind(b.ctx, b.db, &result.UserSlice)
	// var t = result.UserSlice[27]
	// fmt.Println(t.L)
	if err != nil {
		fmt.Println("error %v", err)
		return boadList, err
	}
	return boadList, err
}
