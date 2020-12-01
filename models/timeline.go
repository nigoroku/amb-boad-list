package models

import (
	"time"

	"github.com/volatiletech/null"
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
	CategoryID   null.Int    `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	CategoryName null.String `boil:"category_name" json:"category_name" toml:"category_name" yaml:"category_name"`
	// アクション種類
	ActionType null.String `boil:"action_type" json:"action_type" toml:"action_type" yaml:"action_type"`
}
