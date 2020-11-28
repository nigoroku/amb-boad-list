package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode/utf8"

	// "github.com/kzpolicy/user/models"
	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/net/html/charset"
	"local.packages/models"
)

type ScrapingService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewScrapingService() *ScrapingService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &ScrapingService{ctx, db}
}

func (s *ScrapingService) getPageSummary(url string) models.PageSummary {

	res, _ := http.Get(url)
	if res == nil || res.StatusCode > 400 {
		// 存在しないURLの場合
		return models.PageSummary{}
	}
	fmt.Println(url)
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	fmt.Println(detRslt.Charset)

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	page := models.PageSummary{}
	// title取得
	page.Title = doc.Find("title").Text()
	// image取得
	selection := doc.Find("img")
	image, exists := selection.Attr("src")
	if exists {
		page.ImageUrl = image
	}
	// テキスト取得
	var description string
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.Contains(name, "description") {
			description, _ = s.Attr("content")
		}
	})
	if utf8.RuneCountInString(description) > 300 {
		description = string([]rune(description)[:300])
	}
	page.Description = description + "..."
	page.Url = url
	return page
}
