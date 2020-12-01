package models

import (
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/drivers"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models/generated"
)

type Timeline struct {
	// ユーザー情報
	UserID       int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AccountName  string      `boil:"account_name" json:"account_name" toml:"account_name" yaml:"account_name"`
	AccountImg   []byte      `boil:"account_img" json:"account_img,omitempty" toml:"account_img" yaml:"account_img,omitempty"`
	Introduction null.String `boil:"introduction" json:"introduction,omitempty" toml:"introduction" yaml:"introduction,omitempty"`
	ContentType  null.String `boil:"content_type" json:"content_type,omitempty" toml:"content_type" yaml:"content_type,omitempty"`
	// Output・Input情報
	AchievementID int         `boil:"achievement_id" json:"achievement_id" toml:"achievement_id" yaml:"achievement_id"`
	ReferenceURL  null.String `boil:"reference_url" json:"reference_url,omitempty" toml:"reference_url" yaml:"reference_url,omitempty"`
	Summary       null.String `boil:"summary" json:"summary,omitempty" toml:"summary" yaml:"summary,omitempty"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// カテゴリ
	CategoryID   null.Int              `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	CategoryName null.String           `boil:"category_name" json:"category_name" toml:"category_name" yaml:"category_name"`
	Categories   []generated.MCategory `json:"categories"`
	// アクション種類
	ActionType null.String `boil:"action_type" json:"action_type" toml:"action_type" yaml:"action_type"`
	Lgtm       bool        `json:"lgtm"`
	Stock      bool        `json:"stock"`
	// サマリー
	InputPage  PageSummary `json:"input_page_summary"`
	OutputPage PageSummary `json:"output_page_summary"`
}

var dialect = drivers.Dialect{
	LQ: 0x60,
	RQ: 0x60,

	UseIndexPlaceholders:    false,
	UseLastInsertID:         true,
	UseSchema:               false,
	UseDefaultKeyword:       false,
	UseAutoColumns:          false,
	UseTopClause:            false,
	UseOutputClause:         false,
	UseCaseWhenExistsClause: false,
}

// NewQuery initializes a new Query using the passed in QueryMods
func NewQuery(mods ...qm.QueryMod) *queries.Query {
	q := &queries.Query{}
	queries.SetDialect(q, &dialect)
	qm.Apply(q, mods...)

	return q
}

// TimelineSlice ソート用スライス
type TimelineSlice []*Timeline

func (p TimelineSlice) Len() int {
	return len(p)
}

func (p TimelineSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p TimelineSlice) Less(i, j int) bool {
	return p[i].CreatedAt.After(p[j].CreatedAt)
}
