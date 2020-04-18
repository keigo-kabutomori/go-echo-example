package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// Log :
// サンプル
type Log struct {
	gorm.Model
	Text string `json:"text"`
}

var db *gorm.DB

func main() {
	// 環境変数を読み込む
	loadEnv()

	// データベースに接続
	var err error
	db, err = gorm.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_NAME"))
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	// アプリが終了したらDBと接続解除
	defer db.Close()

	// マイグレーション
	db.AutoMigrate(&Log{})

	// サーバー用のインスタンスの取得
	e := echo.New()

	// ルーティング設定
	e.GET("/logs", getLogs)
	e.GET("/logs/:id", getLog)
	e.POST("/logs", createLog)
	e.PUT("/logs/:id", updateLog)
	e.DELETE("/logs/:id", deleteLog)

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// .envファイルのロード
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ログ一覧を取得
func getLogs(c echo.Context) error {
	// ログ一覧をDBから取得
	var l []Log
	if err := db.Find(&l).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, l)
}

// 特定のログを取得
func getLog(c echo.Context) error {
	// urlからidを取得
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// DBから該当のレコードを取得
	l := new(Log)
	if err := db.First(l, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, l)
}

// ログを作成
func createLog(c echo.Context) error {
	// リクエストパラメータを取得
	l := new(Log)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.Create(&l).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, l)
}

// ログを更新
func updateLog(c echo.Context) error {
	// urlからidを取得
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// DBから該当のレコードを取得
	l := new(Log)
	if err := db.First(l, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// リクエストパラメータを取得
	r := new(Log)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// レコードを更新
	l.Text = r.Text
	if err := db.Save(&l).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, l)
}

// 特定のログを取得
func deleteLog(c echo.Context) error {
	// urlからidを取得
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// DBから該当のレコードを取得
	l := new(Log)
	if err := db.First(l, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 削除
	if err := db.Delete(l).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, l)
}
