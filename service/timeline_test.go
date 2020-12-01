package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"local.packages/models/generated"
)

func Init() {
	// DB接続
	db, err := sql.Open("mysql", "moizumi:base0210@tcp(localhost:3306)/ambitious_test?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}

func MustTx(transactor boil.ContextTransactor, err error) boil.ContextTransactor {
	if err != nil {
		panic(fmt.Sprintf("Cannot create a transactor: %s", err))
	}
	return transactor
}

func TestFindBoadList(t *testing.T) {
	// DB接続
	Init()

	t.Parallel()

	ctx := context.Background()
	db := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = db.Rollback() }()

	// テストデータ用意
	user := &generated.User{Email: "tedt@fwdse.com"}
	user.Insert(ctx, db, boil.Infer())

	to := &generated.Todo{UserID: user.UserID}
	to.Insert(ctx, db, boil.Infer())

	td := &generated.TodoDetail{TodoID: to.TodoID, Content: "content1", Checked: false}
	td.Insert(ctx, db, boil.Infer())

	oa := &generated.OutputAchievement{
		Summary:      null.NewString("Summary1", true),
		ReferenceURL: null.NewString("http://test1.com", true),
		OutputTime:   null.NewString("11:00", true),
		UserID:       user.UserID,
	}
	oa.Insert(ctx, db, boil.Infer())

	ia := &generated.InputAchievement{
		Summary:      null.NewString("Summary2", true),
		ReferenceURL: null.NewString("http://test2.com", true),
		InputTime:    null.NewString("12:00", true),
		UserID:       user.UserID,
	}
	ia.Insert(ctx, db, boil.Infer())

	boadService := &BoadService{ctx, db}
	boadList, err := boadService.FindBoadList()

	assert.NoError(t, err)
	assert.Equal(t, boadList[0].User.Email, user.Email)
	assert.Equal(t, boadList[0].Todos[0].Content, td.Content)
	assert.Equal(t, boadList[0].Output.ReferenceURL, oa.ReferenceURL)
	assert.Equal(t, boadList[0].Input.InputTime, ia.InputTime)
}
