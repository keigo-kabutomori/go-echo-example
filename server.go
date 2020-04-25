package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Log :
// サンプル
type Log struct {
	gorm.Model
	Text string `json:"text" form:"text" query:"text"`
}

// User :
// ログイン用
type User struct {
	gorm.Model
	Email    string `json:"email" form:"email" query:"email" gorm:"unique;primary_key;not null"`
	Password string `json:"password" form:"password" query:"password" gorm:"not null"`
	Token    string `json:"token" form:"token" query:"token"`
}

var db *gorm.DB
var secretKey string

func main() {
	// 環境変数を読み込む
	loadEnv()

	secretKey = os.Getenv("SECRET_KEY")

	var err error
	// データベースに接続
	dbType := os.Getenv("DB_TYPE")
	dbURL := os.Getenv("DATABASE_URL")
	var connection string
	if dbType == "postgres" {
		// heroku
		connection, err = pq.ParseURL(dbURL)
		if err != nil {
			logrus.Fatal(err.Error())
		}
		connection += " sslmode=require"
	} else {
		// local
		connection = dbURL
	}
	logrus.Println("db type:", dbType)
	logrus.Println("db connection:", connection)
	db, err = gorm.Open(dbType, connection)
	if err != nil {
		logrus.Fatal("データベースへの接続に失敗しました")
	}
	// アプリが終了したらDBと接続解除
	defer db.Close()

	// マイグレーション
	db.AutoMigrate(&Log{}, &User{})

	// サーバー用のインスタンスの取得
	e := echo.New()

	e.POST("/signup", signup)
	e.POST("/signin", signin)

	// ここからはログインが必用
	r := e.Group("")
	r.Use(middleware.JWT([]byte(secretKey)))
	// ルーティング設定
	r.GET("/logs", getLogs)
	r.GET("/logs/:id", getLog)
	r.POST("/logs", createLog)
	r.PUT("/logs/:id", updateLog)
	r.DELETE("/logs/:id", deleteLog)

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// ===============================
// COMMON

// .envファイルのロード
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file")
	}
}

// パスワードハッシュを作る
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// パスワードがハッシュにマッチするかどうかを調べる
func passwordVerify(hash string, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

// ===============================
// API
// ----
// Users
// Emailからユーザーを検索
func getUserByEmail(email string) (*User, error) {
	u := new(User)
	if err := db.Where("email = ?", email).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// サインアップ
func signup(c echo.Context) (err error) {
	// フォームから値を取得
	email := c.FormValue("email")
	password := c.FormValue("password")

	// インスタンスの作成
	u := new(User)
	u.Email = email
	// パスワードのハッシュ化
	if u.Password, err = passwordHash(password); err != nil {
		logrus.Warn("signup", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// JWT トークン作成
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DBに保存
	u.Token = t
	if err := db.Create(&u).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 結果を返す
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// ログイン
func signin(c echo.Context) (err error) {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// ユーザーを検索
	u := new(User)
	if u, err = getUserByEmail(email); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// パスワードチェック
	if err = passwordVerify(u.Password, password); err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	// トークンを作成
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// トークンを保存
	u.Token = t
	if err := db.Save(&u).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 結果を返す
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// ----
// Logs
// ログ一覧を取得
func getLogs(c echo.Context) error {
	// ログ一覧をDBから取得
	var l []Log
	if err := db.Find(&l).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, l)
}

// ログを作成
func createLog(c echo.Context) error {
	// リクエストパラメータを取得
	l := new(Log)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := db.Create(&l).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, l)
}

// ログを更新
func updateLog(c echo.Context) error {
	// urlからidを取得
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// DBから該当のレコードを取得
	l := new(Log)
	if err := db.First(l, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// リクエストパラメータを取得
	r := new(Log)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// レコードを更新
	l.Text = r.Text
	if err := db.Save(&l).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, l)
}

// 特定のログを取得
func deleteLog(c echo.Context) error {
	// urlからidを取得
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// DBから該当のレコードを取得
	l := new(Log)
	if err := db.First(l, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 削除
	if err := db.Delete(l).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, l)
}
